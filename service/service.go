package service

import (
	"github.com/liangpengcheng/qcontinuum/network"
)

// Service 服务定义
type Service interface {
	GetName() string
	GetVersion() int32
	GetProcessor() *network.Processor
}
