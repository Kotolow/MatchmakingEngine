package main

import (
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

const GroupSize = 5

func main() {
	r := gin.Default()
	r.POST("/users", usersHandler)
	go func() {
		err := r.Run(":8080")
		if err != nil {
			return
		}
	}()

	var wg sync.WaitGroup
	radius := getRadius(500, 90)

	for {
		wg.Add(1)
		go func() {
			defer wg.Done()
			queue := redisQueue()
			DBSCAN(&queue, radius)
		}()
		wg.Wait()
	}
}

func usersHandler(c *gin.Context) {
	var player Player
	if c.ShouldBindJSON(&player) == nil {
		player.CreationTs = time.Now()
		go redisPush(player, false)
	}
	c.String(200, "Success")
}
