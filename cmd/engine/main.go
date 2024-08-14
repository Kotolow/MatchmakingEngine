package main

import (
	"MatchmakingEngine/internal/api/v1"
	"MatchmakingEngine/internal/config"
	"MatchmakingEngine/internal/repository"
	"MatchmakingEngine/internal/service"
	"MatchmakingEngine/pkg/utils"
	"github.com/gin-gonic/gin"
	"sync"
)

func main() {
	config.ConfigInit()

	r := gin.Default()
	r.POST("/users", v1.UsersHandler)
	go func() {
		err := r.Run(config.AppConfig.Port)
		if err != nil {
			return
		}
	}()

	var wg sync.WaitGroup
	radius := utils.GetRadius(config.AppConfig.MaxSkillDiff, config.AppConfig.MaxSkillDiff)

	for {
		wg.Add(1)
		go func() {
			defer wg.Done()
			queue := repository.RedisQueue()
			service.DBSCAN(&queue, radius)
		}()
		wg.Wait()
	}
}
