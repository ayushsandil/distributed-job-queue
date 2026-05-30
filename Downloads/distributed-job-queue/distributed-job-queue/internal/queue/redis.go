package queue

import (
    "context"
    "os"

    "github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx = context.Background()

func InitRedis() {
    Client = redis.NewClient(&redis.Options{
        Addr: os.Getenv("REDIS_ADDR"),
    })
}

func PushJob(queueName string, jobID string) error {

	return Client.LPush(
		Ctx,
		queueName,
		jobID,
	).Err()
}

func PopJob(queueName string) (string, error) {

	return Client.RPop(
		Ctx,
		queueName,
	).Result()
}
