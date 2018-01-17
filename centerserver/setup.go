package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/liangpengcheng/qcontinuum/network"

	"github.com/liangpengcheng/qcontinuum/base"
)

type config struct {
	RPCPort string
}

var cfg config

type centerServer struct {
	server  *network.Server
	proc    *network.Processor
	gateway *gatewayManger
	servMgr *serviceManager
}

func newCenterServer() *centerServer {
	fp, err := os.Open("../runtime/center.json")
	base.CheckError(err, "open center.json error ")
	buf, err := ioutil.ReadAll(fp)
	base.CheckError(err, "read center.json error ")
	err = json.Unmarshal(buf, &cfg)
	base.CheckError(err, "json unmarshal error ")

	ret := centerServer{}
	ret.server, err = network.NewTCP4Server(cfg.RPCPort)
	base.PanicError(err, "create center server failed ")
	ret.proc = network.NewProcessor()
	ret.proc.AddEventCallback(network.RemoveEvent, ret.onLostConnection)
	ret.gateway = ret.createGatewayManger()
	ret.servMgr = newServiceManager()
	go ret.server.BlockAccept(ret.proc)
	return &ret
}
func (center *centerServer) onLostConnection(e *network.Event) {
	if center.gateway.removeGateway(e.Peer) {
		return
	}
	if center.servMgr.checkLostConnection(e.Peer) {
		return
	}
}
func (center *centerServer) exit(reasion string) {
	center.proc.EventChan <- &network.Event{
		ID:    network.ExitEvent,
		Param: reasion,
	}
}
