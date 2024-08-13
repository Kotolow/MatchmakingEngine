package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

const GroupSize = 5

var groupId = 0

var PlayerQueue Queue

func main() {
	r := gin.Default()
	r.POST("/users", usersHandler)
	go func() {
		r.Run(":8080")
	}()

	var notFinished []Group
	var wg sync.WaitGroup
	radius := getRadius(500, 90)

	for {
		wg.Add(1)
		go func() {
			defer wg.Done()
			notFinished = DBSCAN(&PlayerQueue, radius, notFinished)
			time.Sleep(2 * time.Second)
		}()
		wg.Wait()
	}
}

func usersHandler(c *gin.Context) {
	var player Player
	if c.ShouldBindJSON(&player) == nil {
		player.creationTs = time.Now()
		PlayerQueue.Push(player)
		fmt.Println(player.Name, player.Skill, player.Latency, player.creationTs)
	}

	c.String(200, "Success")
}
