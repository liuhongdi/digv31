package global

import (
	"github.com/go-redis/redis"
)
//保存redis连接
var (
	RedisDb *redis.Client
)

//创建到redis的连接
func SetupRedisDb() (error) {

	RedisDb = redis.NewClient(&redis.Options{
		Addr:     RedisSetting.Addr,
		Password: RedisSetting.Password, // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
