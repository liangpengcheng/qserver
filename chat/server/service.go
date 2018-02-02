package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/liangpengcheng/qcontinuum/base"
	"github.com/liangpengcheng/qserver/chat"
	"github.com/liangpengcheng/qserver/service"
)

type chatService struct {
	handler *service.RPCHandler
	manager *service.Manager
}

func newChatService() *service.Manager {
	handler := service.NewHandler()
	handler.ImmediateMode = true
	serv := &chatService{
		handler: handler,
	}
	handler.AddCallback(int32(chat.C2GIMMessage_ID), serv.onChat)
	manager := service.NewManager(serv, "127.0.0.1:8888", "127.0.0.1:7777")
	serv.manager = manager
	return manager
}
func (cs *chatService) GetName() string {
	return "chat service"
}

func (cs *chatService) GetVersion() int32 {
	return 1
}
func (cs *chatService) GetProcessor() *service.RPCHandler {
	return cs.handler
}

func (cs *chatService) onChat(msg *service.RPCMessage) (proto.Message, int32) {
	chatmsg := chat.C2GIMMessage{}
	err := proto.Unmarshal(msg.MSG, &chatmsg)
	chatresult := chat.G2CIMChatResult{}
	if base.CheckError(err, "unmarshal chat") {
		chatto := chat.G2CIMMessage{}
		chatto.From = msg.UID
		chatto.Content = chatmsg.GetContent()
		chatto.Channel = chatmsg.GetChannel()
		if chatto.Channel&uint32(chat.ChatChannel_USERCHANNEL) > 0 {
			cs.manager.SendMessageTo(chatmsg.GetTouser(), &chatto, int32(chat.G2CIMMessage_ID))
		} else if chatto.Channel&uint32(chat.ChatChannel_WORLDCHANNEL) > 0 {

		}
		chatresult.Result = "Success"
	} else {
		chatresult.Result = "Failed"
	}
	return &chatresult, int32(chat.G2CIMChatResult_ID)
}
