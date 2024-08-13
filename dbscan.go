package main

import "sort"

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

func DBSCAN(queue *Queue, radius float64, notFinished []Group) []Group {
	for gIndex, group := range notFinished {
		sort.Sort(byEDist{queue.Players, group.Players[0]})
		for index, player := range queue.Players {
			if euclideanDistance(player, group.Players[0]) <= radius {
				group.Players = append(group.Players, player)
			} else {
				if group.IsFull() {
					queue.Players = queue.Players[index+1:]
					groupOutput(groupId, group)
					groupId++
					notFinished = append(notFinished[:gIndex], notFinished[gIndex+1:]...) //убираем сформированную группу
				}
				break
			}
		}
	}
	for len(queue.Players) > 0 {
		neighbours := Group{RangeQuery(queue, queue.Pop(), radius)}
		if neighbours.IsFull() {
			groupOutput(groupId, neighbours)
			groupId++
		} else {
			notFinished = append(notFinished, neighbours)
		}
	}
	return notFinished
}
