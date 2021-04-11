package cache

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/liuhongdi/digv31/global"
	"github.com/liuhongdi/digv31/model"
	"time"
)
//token的过期时长
const ArticleDuration = time.Minute * 5

//cache的名字
func getUserCacheName(userToken string) (string) {
	return "user_"+userToken
}

//从cache得到一个user
func GetOneUserCache(userToken string) (*model.User,error) {
	key := getUserCacheName(userToken);
	val, err := global.RedisDb.Get(key).Result()
	if (err == redis.Nil || err != nil) {
		return nil,err
	} else {
		user := model.User{}
		if err := json.Unmarshal([]byte(val), &user); err != nil {
			//t.Error(target)
			return nil,err
		}
		return &user,nil
	}
}
//向cache保存一个user
func SetOneUserCache(userToken string,user *model.User) (error) {
	key := getUserCacheName(userToken);
	content,err := json.Marshal(user)
	if (err != nil){
		fmt.Println(err)
		return err;
	}
	errSet := global.RedisDb.Set(key, content, ArticleDuration).Err()
	if (errSet != nil) {
		return errSet
	}
	return nil
}
//delete
func DelOneUserCache(userToken string) (error) {
	key := getUserCacheName(userToken);
	//global.RedisDb.d
	errSet := global.RedisDb.Del(key).Err()
	if (errSet != nil) {
		return errSet
	}
	return nil
}
