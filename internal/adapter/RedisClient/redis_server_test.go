package RedisClient

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestRedis(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.2.32:6379",
		Password: "yourpassword", // no password set
		DB:       0,              // use default DB
		PoolSize: 10,             // 連接詞數量
	})
	p := NewRedisRepository(client)
	pong, err := p.Client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)
	data := make(map[string]interface{})
	data["test4"] = 1
	data["test5"] = 2
	data["test6"] = 3
	err = p.Client.HSet(context.Background(), "testKey1", data).Err()
	p.Client.HDel(context.Background(), "testKey1", "test", "test2")
	fmt.Println(err)
	field, err := p.Client.HGetAll(context.Background(), "testKey1").Result()
	fmt.Println(err)
	fmt.Println(field)
}
