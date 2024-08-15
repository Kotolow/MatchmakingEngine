package models

import (
	"MatchmakingEngine/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEuclideanDistance(t *testing.T) {
	player1 := Player{
		Name:    "Player1",
		Skill:   1339,
		Latency: 45,
	}
	player2 := Player{
		Name:    "Player2",
		Skill:   1312,
		Latency: 6,
	}

	expectedDistance := 47.4341649025256
	actualDistance := player1.EuclideanDistance(player2)

	assert.InDelta(t, expectedDistance, actualDistance, 1e-9, "EuclideanDistance should be calculated correctly")
}

func TestIsFull(t *testing.T) {
	config.AppConfig = config.Config{
		GroupSize: 3,
	}

	group := Group{
		Players: []Player{
			{Name: "Player1"},
			{Name: "Player2"},
		},
	}

	assert.False(t, group.IsFull(), "IsFull should return false when group is not full")

	group.Players = append(group.Players, Player{Name: "Player3"})

	assert.True(t, group.IsFull(), "IsFull should return true when group is full")
}
