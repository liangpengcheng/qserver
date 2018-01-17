package main

import (
	"net"
	"strings"
	"sync"

	"github.com/liangpengcheng/qcontinuum/base"
	"github.com/liangpengcheng/qcontinuum/network"
	"github.com/liangpengcheng/qserver/protocol"
)

type gateway struct {
	centerPeer *network.ClientPeer
	proc       *network.Processor
	server     *network.Server
	usermap    sync.Map
}

func (g *gateway) getUser(uid uint64) *client {
	if v, ok := g.usermap.Load(uid); ok {
		return v.(*client)
	}
	return nil
}
func (g *gateway) addUser(uid uint64, c *client) {
	if v, ok := g.usermap.Load(uid); ok {
		c := v.(*client)
		c.connection.Connection.Close()
	}
	g.usermap.Store(uid, c)
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
