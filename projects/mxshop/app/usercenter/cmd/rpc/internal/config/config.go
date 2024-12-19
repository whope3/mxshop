package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MysqlConf MysqlConf
	RedisConf redis.RedisConf
}

type MysqlConf struct {
	Uri string
}
