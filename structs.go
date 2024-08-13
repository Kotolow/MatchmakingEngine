package main

import "time"

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
