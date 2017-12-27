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
}

func newCenterServer() *centerServer {
	fp, err := os.Open("../center.json")
	base.CheckError(err, "open center.json error ")
	buf, err := ioutil.ReadAll(fp)
	base.CheckError(err, "read center.json error ")
	err = json.Unmarshal(buf, &cfg)
	base.CheckError(err, "json unmarshal error ")

	ret := centerServer{}
	ret.server, err = network.NewTCP4Server(cfg.RPCPort)
	base.PanicError(err, "create center server failed ")
	ret.proc = network.NewProcessor()
	ret.gateway = ret.createGatewayManger()
	go ret.server.BlockAccept(ret.proc)
	return &ret
}
func (center *centerServer) exit(reasion string) {
	center.proc.EventChan <- &network.Event{
		ID:    network.ExitEvent,
		Param: reasion,
	}
}
