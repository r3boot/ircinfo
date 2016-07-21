package proto

import (
	"gopkg.in/redis.v3"
)

var Redis *redis.Client

func NewRedisClient(addr string, pass string, db int64) (rc *redis.Client) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})
	return
}

func HLHas(list_name string, key string) bool {
	return Redis.HMGet(list_name, key).Val() != nil
}

func HLAdd(list_name string, key string, value string) bool {
	if HLHas(list_name, key) {
		return false
	}

	if _, err := Redis.HSet(list_name, key, value).Result(); err != nil {
		Log.Warning("Failed to add entry to " + list_name + ": " + err.Error())
		return false
	}
	return true
}

func LHas(list_name string, value string) bool {
	llen := Redis.LLen(list_name).Val()
	for _, item := range Redis.LRange(list_name, 0, llen).Val() {
		if item == value {
			return true
		}
	}

	return false
}

func LAdd(list_name string, value string) {
	Redis.LPush(list_name, value)
}
