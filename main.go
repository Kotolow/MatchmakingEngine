package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

const GroupSize = 5

var groupId = 0

type Player struct {
	Name       string  `json:"name"`
	Skill      float64 `json:"skill"`
	Latency    float64 `json:"latency"`
	creationTs time.Time
}

type Group struct {
	Players []Player
}

func (g *Group) IsFull() bool {
	if len(g.Players) == GroupSize {
		return true
	}
	return false
}

func euclideanDistance(p1, p2 Player) float64 {
	return math.Sqrt(math.Pow(p1.Skill-p2.Skill, 2) + math.Pow(p1.Latency-p2.Latency, 2))
}

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

func main() {
	rand.Seed(time.Now().UnixNano())

	players := make([]Player, 100)

	for i := range players {
		players[i] = Player{
			Name:       fmt.Sprintf("%dE", i+1),
			Skill:      rand.Float64() * 10000,
			Latency:    rand.Float64() * 100,
			creationTs: time.Now(),
		}
	}

	queue := Queue{players}
	sort.Sort(byEDist{queue.Players, players[0]})

	radius := getRadius(500, 90)
	var notFinished []Group
	notFinished = DBSCAN(&queue, radius, notFinished)
	var temp int
	for {
		players = make([]Player, rand.Intn(1000-100)+100)

		for i := range players {
			players[i] = Player{
				Name:       fmt.Sprintf("%dF", temp+1),
				Skill:      rand.Float64() * 10000,
				Latency:    rand.Float64() * 100,
				creationTs: time.Now(),
			}
			temp++
		}
		queue = Queue{players}
		notFinished = DBSCAN(&queue, radius, notFinished)
		time.Sleep(10 * time.Second)
	}
}
