package main

import "github.com/liangpengcheng/qserver/service"

type chatService struct {
	handler *service.RPCHandler
}

func newChatService() *service.Manager {
	handler := service.NewHandler()
	serv := &chatService{
		handler: handler,
	}
	manager := service.NewManager(serv, "127.0.0.1:8888", "127.0.0.1:7777")
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
