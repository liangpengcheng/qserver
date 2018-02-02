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

var ucount = 0

func (g *gateway) onLogin(msg *network.Message) {

	uid := uint64(ucount)
	ucount++
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
	if u, ok := g.connectionMap.Load(msg.Peer); ok {
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
				if resp != nil && resp.GetBroadcastThisGateway() != nil {
					//广播
					g.broadcast(resp.GetBroadcastThisGateway())
				}
			}
		}
	} else {
		base.LogInfo("can't find user kick connection")
		if msg.Peer.Connection != nil {
			msg.Peer.Connection.Close()
		}
	}
}

func (g *gateway) broadcast(buf []byte) {
	g.userMap.Range(
		func(key, value interface{}) bool {
			c := value.(*client)
			c.connection.SendMessageBuffer(buf)
			return true
		})
}
