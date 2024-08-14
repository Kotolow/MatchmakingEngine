package models

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

type ByEDist struct {
	Players    []Player
	CenteredEl Player
}

func (q ByEDist) Len() int {
	return len(q.Players)
}

func (q ByEDist) Swap(i, j int) {
	q.Players[i], q.Players[j] = q.Players[j], q.Players[i]
}
func (q ByEDist) Less(i, j int) bool {
	return q.CenteredEl.EuclideanDistance(q.Players[i]) < q.CenteredEl.EuclideanDistance(q.Players[j])
}
