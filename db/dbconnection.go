package db

import (
	"os"

	"github.com/edwinnduti/natschat/logger"
	"github.com/go-redis/redis"
)

func ConnectDb() (*redis.Client, error) {
	// connect to redis
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL_ADDRESS"), // host:port of the redis server
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// ping redis
	pong, err := client.Ping().Result()
	if err != nil{
		logger.Error.Println("Cannot connect to redis", err)
		return nil, err
	}
	// log redis connection
	logger.Info.Println("Success connection to redis db: ", pong)
	// return redis client
	return client, nil
}