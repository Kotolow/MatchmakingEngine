package main

type Queue struct {
	Players []Player
}

func (q *Queue) Push(player Player) {
	q.Players = append(q.Players, player)
}

func (q *Queue) Pop() Player {
	player := q.Players[0]
	q.Players = q.Players[1:]
	return player
}

type byEDist struct {
	players    []Player
	centeredEl Player
}

func (q byEDist) Len() int {
	return len(q.players)
}

func (q byEDist) Swap(i, j int) {
	q.players[i], q.players[j] = q.players[j], q.players[i]
}
func (q byEDist) Less(i, j int) bool {
	return euclideanDistance(q.centeredEl, q.players[i]) < euclideanDistance(q.centeredEl, q.players[j])
}
