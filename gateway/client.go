package main

import (
	"github.com/liangpengcheng/qcontinuum/base"
	"github.com/liangpengcheng/qserver/protocol"
	//	"github.com/liangpengcheng/qcontinuum/db"
	"github.com/liangpengcheng/qcontinuum/network"
)

type client struct {
	connection *network.ClientPeer
}

func (g *gateway) onLogin(msg *network.Message) {
	c := &client{
		connection: msg.Peer,
	}
	uid := uint64(0)
	g.addUser(uid, c)
	logres := protocol.G2CTokenLoginResult{
		Result: "Success",
		Uid:    uid,
	}
	msg.Peer.SendMessage(&logres, int32(protocol.G2CTokenLoginResult_ID))
}

func (g *gateway) onUnhandledMsg(msg *network.Message) {
	base.LogDebug("incomming unhandled messsage :%d", msg.Head.ID)
}
