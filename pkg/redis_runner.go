package pkg

import "github.com/go-redis/redis/v8"

type RedisRunner struct {
	client *redis.Client
}

func (r *RedisRunner) Query(s *string, args ...any) {

}

func NewRedisRunner(client *redis.Client) *RedisRunner {
	return &RedisRunner{
		client: client,
	}
}
