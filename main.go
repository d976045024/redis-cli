package main

import (
	"log"

	"github.com/redis-cli/pkg"
	"github.com/redis-cli/pkg/types"
)

func main() {
	opts := map[string]any{
		types.Host: "127.0.0.1",
		types.Port: 6379,
	}
	r, err := pkg.NewRedisRunner(opts)
	if err != nil {
		panic(err)
	}
	err = r.TestConnect()
	if err != nil {
		return
	} else {
		log.Println("redis connect success")
	}
}
