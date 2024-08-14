package service

import (
	"MatchmakingEngine/internal/config"
	"MatchmakingEngine/internal/models"
	"MatchmakingEngine/internal/repository"
	"sort"
)

func RangeQuery(queue *models.Queue, player models.Player, radius float64) []models.Player {
	var neighbours []models.Player
	neighbours = append(neighbours, player)
	if len(queue.Players) == 0 {
		return neighbours
	}
	for len(queue.Players) > 0 && queue.Players[0].EuclideanDistance(player) <= radius && len(neighbours) < config.AppConfig.GroupSize {
		sort.Sort(models.ByEDist{Players: queue.Players, CenteredEl: queue.Players[0]})
		el := queue.Pop()
		neighbours = append(neighbours, el)
	}
	return neighbours
}

func DBSCAN(queue *models.Queue, radius float64) {
	for len(queue.Players) > 0 {
		neighbours := models.Group{Players: RangeQuery(queue, queue.Pop(), radius)}
		if neighbours.IsFull() {
			groupId := repository.RedisGetId()
			neighbours.GroupOutput(groupId)
			repository.RedisSetId(groupId + 1)
		} else {
			for _, player := range neighbours.Players {
				go repository.RedisPush(player, true)
			}
		}
	}
}
