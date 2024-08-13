package main

import (
	"encoding/json"
	"time"
)

type Player struct {
	Name       string    `json:"name"`
	Skill      float64   `json:"skill"`
	Latency    float64   `json:"latency"`
	CreationTs time.Time `json:"ts"`
}

func (p Player) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p Player) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &p)
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
