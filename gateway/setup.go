package main

import (
	"net"
	"strings"
	"sync"

	"github.com/liangpengcheng/qcontinuum/base"
	"github.com/liangpengcheng/qcontinuum/network"
	"github.com/liangpengcheng/qserver/protocol"
	"google.golang.org/grpc"
)

type gateway struct {
	centerPeer    *network.ClientPeer
	proc          *network.Processor
	server        *network.Server
	userMap       sync.Map
	connectionMap sync.Map
	serviceMap    sync.Map
}

func (g *gateway) getUser(uid uint64) *client {
	if v, ok := g.userMap.Load(uid); ok {
		return v.(*client)
	}
	return nil
}
func (g *gateway) removeUser(uid uint64) {
	if v, ok := g.userMap.Load(uid); ok {
		c := v.(*client)
		g.userMap.Delete(uid)
		g.connectionMap.Delete(c.connection)
	}
}
func (g *gateway) addUser(uid uint64, c *client) {
	if v, ok := g.userMap.Load(uid); ok {
		c := v.(*client)
		c.connection.Connection.Close()
	}
	g.userMap.Store(uid, c)
	g.connectionMap.Store(c.connection, c)
	base.LogDebug("user(%d) added", uid)
}
func (g *gateway) addService(serv *protocol.Service) {
	base.LogDebug("dialing rpc :%s", serv.Address)
	opts := grpc.WithInsecure()
	// 连接rpc
	conn, err := grpc.Dial(serv.Address, opts)

	// 连接失败
	if !base.CheckError(err, "dial rpc server ") {
		return
	}
	rpc := protocol.NewGatewayCallServiceClient(conn)
	thisservice := service{
		serv:    serv,
		rpc:     rpc,
		rpcConn: conn,
	}
	for _, v := range serv.GetHandler() {
		if sl, ok := g.serviceMap.Load(v); ok {
			servL := sl.([]*service)
			servL = append(servL, &thisservice)
			g.serviceMap.Store(v, servL)
		} else {
			g.serviceMap.Store(v, []*service{&thisservice})
		}
	}
}
func (g *gateway) removeService(serv *protocol.Service) {
	for _, v := range serv.GetHandler() {
		if sl, ok := g.serviceMap.Load(v); ok {
			servL := sl.([]*service)
			for n, v := range servL {
				if v.serv.Address == serv.Address {
					v.onDestroy()
					servL = append(servL[:n], servL[n+1:]...)
					g.serviceMap.Store(v, servL)
					break
				}
			}
		}
	}
}
func newGateway() *gateway {
	connection, err := net.Dial("tcp", cfg.Center)
	base.PanicError(err, "dial center server")
	processor := network.NewProcessor()
	processor.ImmediateMode = true
	g := gateway{
		centerPeer: &network.ClientPeer{
			Connection: connection,
			Proc:       processor,
		},
		proc: processor,
	}
	g.startListen()
	g.regMsgCallback()
	regGateway := protocol.G2CRegisterGateway{
		Gate: &protocol.Gateway{
			Id:      cfg.ID,
			Address: cfg.Address,
		},
	}
	processor.AddCallback(int32(protocol.C2GTokenLogin_ID), g.onLogin)
	processor.UnHandledHandler = g.onUnhandledMsg
	go g.centerPeer.ConnectionHandler()
	go g.centerPeer.SendMessage(&regGateway, int32(protocol.G2CRegisterGateway_ID))
	return &g
}
func (g *gateway) startListen() {
	addr := strings.Split(cfg.Address, ":")
	var port string
	if len(addr) == 2 {
		port = ":" + addr[1]
	}
	if len(addr) == 1 {
		port = ":" + addr[0]
	}
	var err error
	g.server, err = network.NewTCP4Server(port)
	base.PanicError(err, "gateway listen")
	go g.server.BlockAccept(g.proc)
}
func (g *gateway) exit(reasion string) {
	g.proc.EventChan <- &network.Event{
		ID:    network.ExitEvent,
		Param: reasion,
	}
}

func (g *gateway) onLostConnection(e *network.Event) {
	if v, ok := g.connectionMap.Load(e.Peer); ok {
		c := v.(*client)
		g.removeUser(c.uid)
		return
	}

}
