package cache

import (
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
)

const RandomUidKey = "random:uid"

type RedisModel interface {
	Get(key string) (string, error)
	IncrBy(key string, increment int64) (int64, error)
}

type redisModel struct {
	redis      *redis.Redis
	localCache *collection.Cache
}

func (m redisModel) Get(key string) (string, error) {
	return m.redis.Get(key)
}

func (m redisModel) IncrBy(key string, increment int64) (int64, error) {
	return m.redis.Incrby(key, increment)
}

func NewRedisModel(conf redis.RedisConf) (RedisModel, error) {
	r, err := redis.NewRedis(conf)
	if err != nil {
		return nil, err
	}
	cache, err := collection.NewCache(time.Minute)
	if err != nil {
		return nil, err
	}
	return redisModel{
		redis:      r,
		localCache: cache,
	}, nil
}
