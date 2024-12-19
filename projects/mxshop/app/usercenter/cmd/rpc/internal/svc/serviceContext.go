package svc

import (
	"mxshop/app/usercenter/cmd/rpc/internal/config"
	"mxshop/app/usercenter/model/cache"
	"mxshop/app/usercenter/model/mysql"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  mysql.UserModel
	RedisModel cache.RedisModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	userModel := mysql.NewUserModel(c.MysqlConf.Uri)
	redisModel, err := cache.NewRedisModel(c.RedisConf)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:     c,
		UserModel:  userModel,
		RedisModel: redisModel,
	}
}
