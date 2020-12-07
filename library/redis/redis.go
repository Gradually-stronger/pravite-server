package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	once sync.Once
)

// Init 初始化redis客户端
func Init(ctx context.Context, addr, password string, db int) *redis.Client {
	var internalClient *redis.Client
	once.Do(func() {
		internalClient = newCli(ctx, addr, password, db)
	})
	return internalClient
}

// New 创建redis客户端实例
func newCli(ctx context.Context, addr, password string, db int) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	cmd := cli.Ping(ctx)
	if err := cmd.Err(); err != nil {
		panic(err)
	}

	return cli
}
