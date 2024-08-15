package models

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestQueue_Push(t *testing.T) {
	q := Queue{}

	player1 := Player{Name: "Player1"}
	player2 := Player{Name: "Player2"}

	q.Push(player1)
	assert.Equal(t, 1, len(q.Players), "Queue should have 1 player after first push")
	assert.Equal(t, player1, q.Players[0], "First player should be Player1")

	q.Push(player2)
	assert.Equal(t, 2, len(q.Players), "Queue should have 2 players after second push")
	assert.Equal(t, player2, q.Players[1], "Second player should be Player2")
}

func TestQueue_Pop(t *testing.T) {
	q := Queue{}

	player1 := Player{Name: "Player1"}
	player2 := Player{Name: "Player2"}

	q.Push(player1)
	q.Push(player2)

	poppedPlayer := q.Pop()
	assert.Equal(t, player1, poppedPlayer, "First player popped should be Player1")
	assert.Equal(t, 1, len(q.Players), "Queue should have 1 player left after pop")
	assert.Equal(t, player2, q.Players[0], "Remaining player should be Player2")

	poppedPlayer = q.Pop()
	assert.Equal(t, player2, poppedPlayer, "Second player popped should be Player2")
	assert.Equal(t, 0, len(q.Players), "Queue should be empty after second pop")
}

func TestByEDist_Sort(t *testing.T) {
	centerPlayer := Player{Name: "Center", Skill: 1339.1648813757522, Latency: 45.10202118006291}

	player1 := Player{Name: "Player1", Skill: 1977.0730531810498, Latency: 22.265818907035776}
	player2 := Player{Name: "Player2", Skill: 1371.8532835597957, Latency: 33.62658235583313}
	player3 := Player{Name: "Player3", Skill: 1312.9516520245197, Latency: 6.36711293399273}

	players := []Player{player1, player2, player3}
	sortedPlayers := []Player{player2, player3, player1}

	byEDist := ByEDist{Players: players, CenteredEl: centerPlayer}
	sort.Sort(byEDist)

	assert.Equal(t, sortedPlayers, byEDist.Players, "Players should be sorted by Euclidean distance from center player")
}
