package redisdemo

import (
	"github.com/go-redis/redis"
)

var Rdb *redis.Client

func init() {

	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
