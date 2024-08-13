package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var ctx = context.Background()

func redisPush(player Player, reversed bool) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
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

func redisPop(client *redis.Client, player *Player) {
	result, _ := client.LPop(ctx, "queue").Result()
	err := json.Unmarshal([]byte(result), player)
	if err != nil {
		fmt.Printf("error unmarshalling player: %w", err)
	}
}

func redisQueue() Queue {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	var queue Queue
	var player Player

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

func redisGetId() int {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	groupId, err := client.Get(ctx, "groupId").Result()
	if errors.Is(err, redis.Nil) {
		redisSetId(1)
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

func redisSetId(value int) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	err := client.Set(ctx, "groupId", value, 0).Err()
	if err != nil {
		fmt.Printf("error setting groupId: %w", err)
	}
}
