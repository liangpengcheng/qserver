package main

import (
	"context"

	"github.com/liangpengcheng/qcontinuum/base"
	"github.com/liangpengcheng/qserver/protocol"
	//	"github.com/liangpengcheng/qcontinuum/db"
	"github.com/liangpengcheng/qcontinuum/network"
)

type client struct {
	connection *network.ClientPeer
	uid        uint64
}

func (g *gateway) onLogin(msg *network.Message) {
	uid := uint64(0)
	c := &client{
		connection: msg.Peer,
		uid:        uid,
	}
	g.addUser(uid, c)
	logres := protocol.G2CTokenLoginResult{
		Result: "Success",
		Uid:    uid,
	}
	go msg.Peer.SendMessage(&logres, int32(protocol.G2CTokenLoginResult_ID))
}

// rpc调用service，并且将结果返回给客户端
func (g *gateway) onUnhandledMsg(msg *network.Message) {
	base.LogDebug("incomming unhandled messsage :%d", msg.Head.ID)
	if u, ok := g.userMap.Load(msg.Peer); ok {
		if s, ok := g.serviceMap.Load(msg.Head.ID); ok {
			user := u.(*client)
			sl := s.([]*service)
			idx := user.uid % uint64(len(sl))
			req := protocol.RPCRequest{
				Userid:    user.uid,
				Messageid: msg.Head.ID,
				Body:      msg.Body,
			}
			resp, err := sl[idx].rpc.Request(context.Background(), &req)
			if base.CheckError(err, "rpc call response :") {
				if resp != nil && resp.GetResponse() != nil {
					msg.Peer.SendMessageBuffer(resp.GetResponse())
				}
			}
		}
	}
}
