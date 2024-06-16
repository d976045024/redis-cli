package pkg

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/redis-cli/pkg/constants"
	"github.com/redis-cli/pkg/models"
	"github.com/redis-cli/pkg/types"
)

type RedisRunner struct {
	client *redis.Client
}

func (r *RedisRunner) TestConnect() error {
	return r.client.Ping(context.Background()).Err()
}

func (r *RedisRunner) Query(ctx context.Context, s *string, args ...any) (*models.QueryResult, error) {
	result := &models.QueryResult{}
	switch *s {
	case constants.REDIS_GET:
		if len(args) != 1 {
			result.Status = constants.REDIS_STATUS_ERROR
			result.Message = fmt.Sprintf("get method requires exactly one argument, but get %d instead\n", len(args))
			return result, errors.New(result.Message)
		}
		key, ok := args[0].(string)
		if !ok {
			result.Status = constants.REDIS_STATUS_ERROR
			result.Message = fmt.Sprintf("set argument must be a string, but get type %T", args[0])
			return result, errors.New(result.Message)
		}
		val, err := r.client.Get(ctx, key).Result()
		if err != nil {
			result.Status = constants.REDIS_STATUS_ERROR
			result.Message = err.Error()
			return result, err
		}
		result.Status = constants.REDIS_STATUS_SUCCESS
		result.Data = val
		return result, nil
	case constants.REDIS_SET:
		if len(args) != 2 {
			result.Status = constants.REDIS_STATUS_ERROR
			result.Message = fmt.Sprintf("get method requires exactly one argument, but get %d instead\n", len(args))
			return result, errors.New(result.Message)
		}
		key, ok := args[0].(string)
		if !ok {
			result.Status = constants.REDIS_STATUS_ERROR
			result.Message = fmt.Sprintf("the first argument of set must be a string, but get type %T", args[0])
			return result, errors.New(result.Message)
		}
		val := args[1]
		ans, err := r.client.Set(ctx, key, val, 0).Result()
		if err != nil {
			return nil, err
		}
		result.Status = constants.REDIS_STATUS_SUCCESS
		result.Data = ans
	}

	return result, nil
}

func NewRedisRunner(opts map[string]any) (*RedisRunner, error) {
	options, err := getOptionFromInput(opts)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(options)
	return &RedisRunner{
		client: client,
	}, nil
}

func getOptionFromInput(opts map[string]any) (*redis.Options, error) {
	var (
		err  error
		host string
		port int
		db   int
	)
	if host, err = getHostFromInput(opts); err != nil {
		return nil, err
	}
	if port, err = getPortFromInput(opts); err != nil {
		return nil, err
	}
	if db, err = getDBFromInput(opts); err != nil {
		return nil, err
	}

	options := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
		DB:   db,
	}

	return options, nil
}

func getHostFromInput(opts map[string]any) (string, error) {
	if _, ok := opts[types.Host]; !ok {
		return constants.DEFAULT_REDIS_HOST, nil
	}
	val, ok := opts[types.Host].(string)
	if !ok {
		return "", fmt.Errorf("variable %v cannot be asserted as type string", opts[types.Host])
	}
	return val, nil
}

func getPortFromInput(opts map[string]any) (int, error) {
	if _, ok := opts[types.Port]; !ok {
		return constants.DEFAULT_REDIS_PORT, nil
	}
	val, ok := opts[types.Port].(int)
	if !ok {
		return 0, fmt.Errorf("variable %v cannot be asserted as type int", opts[types.Port])
	}
	return val, nil
}

func getDBFromInput(opts map[string]any) (int, error) {
	if _, ok := opts[types.DataBase]; !ok {
		return constants.DEFAULT_DATABASE, nil
	}
	val, ok := opts[types.DataBase].(int)
	if !ok {
		return 0, fmt.Errorf("variable %v cannot be asserted as type int", opts[types.DataBase])
	}
	return val, nil
}
