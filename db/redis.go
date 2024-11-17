package db

import (
	"log"
	"sky-take-out-go/utils"

	"github.com/redis/go-redis/v9"
)


func InitRedis() (*redis.Client, error) {
	// init redis
	config, err := utils.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Panic("Redis Load Config Error: ", err)
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: config.Redis.Addr,
		Password: config.Redis.Password,
		DB: config.Redis.DB,
	})

	return rdb, nil
}