package v1

import (
	"MatchmakingEngine/internal/models"
	"MatchmakingEngine/internal/repository"
	"github.com/gin-gonic/gin"
	"time"
)

func UsersHandler(c *gin.Context) {
	var player models.Player
	if c.ShouldBindJSON(&player) == nil {
		player.CreationTs = time.Now()
		go repository.RedisPush(player, false)
	}
	c.String(200, "Success")
}
