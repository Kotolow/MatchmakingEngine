package repository

import (
	"MatchmakingEngine/internal/config"
	"MatchmakingEngine/internal/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var ctx = context.Background()

func RedisPush(player models.Player, reversed bool) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisAddr,
		Password: config.AppConfig.RedisPW,
		DB:       config.AppConfig.RedisDB,
	})
	defer client.Close()
	if reversed {
		_, err := client.LPush(ctx, "queue", player).Result()
		if err != nil {
			fmt.Printf("error pushing to queue: %w", err)
		}
	} else {
		_, err := client.RPush(ctx, "queue", player).Result()
		if err != nil {
			fmt.Printf("error pushing to queue: %w", err)
		}
	}

}

func redisPop(client *redis.Client, player *models.Player) {
	result, _ := client.LPop(ctx, "queue").Result()
	err := json.Unmarshal([]byte(result), player)
	if err != nil {
		fmt.Printf("error unmarshalling player: %w", err)
	}
}

func RedisQueue() models.Queue {
	client := redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisAddr,
		Password: config.AppConfig.RedisPW,
		DB:       config.AppConfig.RedisDB,
	})
	defer client.Close()

	var queue models.Queue
	var player models.Player

	queueLen := redisQueueLen(client)
	for i := 0; i < queueLen; i++ {
		redisPop(client, &player)
		queue.Push(player)
	}

	return queue
}

func redisQueueLen(client *redis.Client) int {
	result, _ := client.LLen(ctx, "queue").Result()
	return int(result)
}

func RedisGetId() int {
	client := redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisAddr,
		Password: config.AppConfig.RedisPW,
		DB:       config.AppConfig.RedisDB,
	})
	defer client.Close()

	groupId, err := client.Get(ctx, "groupId").Result()
	if errors.Is(err, redis.Nil) {
		RedisSetId(1)
		return 1
	} else if err != nil {
		fmt.Printf("error getting groupId: %w", err)
	} else {
		id, err := strconv.Atoi(groupId)
		if err != nil {
			fmt.Printf("error converting groupId: %w", err)
		}
		return id
	}
	return 0
}

func RedisSetId(value int) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisAddr,
		Password: config.AppConfig.RedisPW,
		DB:       config.AppConfig.RedisDB,
	})
	defer client.Close()

	err := client.Set(ctx, "groupId", value, 0).Err()
	if err != nil {
		fmt.Printf("error setting groupId: %w", err)
	}
}
