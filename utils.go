package main

import (
	"fmt"
	"math"
	"time"
)

func getRadius(sDiff, lDiff float64) float64 {
	return math.Sqrt(math.Pow(sDiff, 2) + math.Pow(lDiff, 2))
}

func euclideanDistance(p1, p2 Player) float64 {
	return math.Sqrt(math.Pow(p1.Skill-p2.Skill, 2) + math.Pow(p1.Latency-p2.Latency, 2))
}

func groupOutput(groupId int, group Group) {
	var minSkill, maxSkill, totalSkill float64
	var minLatency, maxLatency, totalLatency float64
	var minTimeSpent, maxTimeSpent, totalTimeSpent time.Duration

	minSkill = group.Players[0].Skill
	maxSkill = group.Players[0].Skill
	minLatency = group.Players[0].Latency
	maxLatency = group.Players[0].Latency
	minTimeSpent = time.Since(group.Players[0].CreationTs)
	maxTimeSpent = time.Since(group.Players[0].CreationTs)

	for _, player := range group.Players {
		if player.Skill < minSkill {
			minSkill = player.Skill
		}
		if player.Skill > maxSkill {
			maxSkill = player.Skill
		}
		totalSkill += player.Skill

		if player.Latency < minLatency {
			minLatency = player.Latency
		}
		if player.Latency > maxLatency {
			maxLatency = player.Latency
		}
		totalLatency += player.Latency

		timeSpent := time.Since(player.CreationTs)
		if timeSpent < minTimeSpent {
			minTimeSpent = timeSpent
		}
		if timeSpent > maxTimeSpent {
			maxTimeSpent = timeSpent
		}
		totalTimeSpent += timeSpent
	}

	avgSkill := totalSkill / float64(len(group.Players))
	avgLatency := totalLatency / float64(len(group.Players))
	avgTimeSpent := totalTimeSpent / time.Duration(len(group.Players))

	fmt.Printf("Group #%d\n", groupId)
	fmt.Printf("Skill: Min = %.2f, Max = %.2f, Avg = %.2f\n", minSkill, maxSkill, avgSkill)
	fmt.Printf("Latency: Min = %.2f ms, Max = %.2f ms, Avg = %.2f ms\n", minLatency, maxLatency, avgLatency)
	fmt.Printf("Time spent in queue: Min = %s, Max = %s, Avg = %s\n", minTimeSpent, maxTimeSpent, avgTimeSpent)
	fmt.Println("Players in group:")

	for _, player := range group.Players {
		fmt.Printf("- %s\n", player.Name)
	}
}
