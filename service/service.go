package service

// Service 服务定义
type Service interface {
	GetName() string
	GetVersion() int32
	GetProcessor() *RPCHandler
}
