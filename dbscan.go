package main

import (
	"sort"
)

func RangeQuery(queue *Queue, player Player, radius float64) []Player {
	var neighbours []Player
	neighbours = append(neighbours, player)
	if len(queue.Players) == 0 {
		return neighbours
	}
	for len(queue.Players) > 0 && euclideanDistance(queue.Players[0], player) <= radius && len(neighbours) < GroupSize {
		sort.Sort(byEDist{queue.Players, queue.Players[0]})
		el := queue.Pop()
		neighbours = append(neighbours, el)
	}
	return neighbours
}

func DBSCAN(queue *Queue, radius float64) {
	for len(queue.Players) > 0 {
		neighbours := Group{RangeQuery(queue, queue.Pop(), radius)}
		if neighbours.IsFull() {
			groupId := redisGetId()
			groupOutput(groupId, neighbours)
			redisSetId(groupId + 1)
		} else {
			for _, player := range neighbours.Players {
				go redisPush(player, true)
			}
		}
	}
}
