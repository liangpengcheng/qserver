package service

import (
	"context"
	"errors"

	"github.com/liangpengcheng/qcontinuum/base"
	"github.com/liangpengcheng/qcontinuum/network"

	"github.com/golang/protobuf/proto"

	"github.com/liangpengcheng/qserver/protocol"
)

type rpcHandler struct {
	Serv Service
}

func (rpc *rpcHandler) Request(ctx context.Context, in *protocol.RPCRequest) (*protocol.RPCResponse, error) {

	msg := &RPCMessage{
		UID:   in.GetUserid(),
		MSGID: in.GetMessageid(),
		MSG:   in.GetBody(),
	}

	proc := rpc.Serv.GetProcessor()
	var ret proto.Message
	var msgid int32
	var ret2 proto.Message
	var msgid2 int32
	var err error
	if proc.ImmediateMode {
		if cb, ok := proc.CallbackMap[in.GetMessageid()]; ok {
			ret, msgid, ret2, msgid2 = cb(msg)
		} else {
			ret = nil
			err = errors.New("can't find rpc callback")
		}
	} else {
		type msgret struct {
			msg proto.Message
			err error
			id  int32

			msg2 proto.Message
			id2  int32
		}
		retchan := make(chan msgret, 1)
		proc.FuncChan <- func() {
			if cb, ok := proc.CallbackMap[in.GetMessageid()]; ok {
				msg, id, msg2, id2 := cb(msg)
				retchan <- msgret{
					msg:  msg,
					id:   id,
					err:  nil,
					msg2: msg2,
					id2:  id2,
				}

			} else {
				retchan <- msgret{
					msg: nil,
					err: errors.New("can't find callback"),
				}
			}
		}
		retmsg := <-retchan
		ret = retmsg.msg
		err = retmsg.err
		ret2 = retmsg.msg2
		msgid2 = retmsg.id2
		msgid = retmsg.id
	}
	resp := protocol.RPCResponse{}
	if ret != nil && err == nil {
		buf, err := network.GetMessageBuffer(ret, msgid)
		if base.CheckError(err, "get message buffer") {
			resp.Response = buf
		}
	}
	if ret2 != nil && err == nil {
		buf, err := network.GetMessageBuffer(ret2, msgid2)
		if base.CheckError(err, "get message buffer for gateway") {
			resp.BroadcastThisGateway = buf
		}
	}
	return &resp, err
}
