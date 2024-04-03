package svc

import (
	"zerodemo/user/rpc/userclient"
	"zerodemo/video/api/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
