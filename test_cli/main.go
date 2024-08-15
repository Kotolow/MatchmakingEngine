package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Player struct {
	Name    string  `json:"name"`
	Skill   float64 `json:"skill"`
	Latency float64 `json:"latency"`
}

func sendRequest(player Player) error {
	url := "http://172.20.0.1:8080/users"

	jsonData, err := json.Marshal(player)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	playerNum := 0
	for {
		numRequests := rand.Intn(10) + 1

		for i := 0; i < numRequests; i++ {
			player := Player{
				Name:    fmt.Sprintf("testPlayer%d", playerNum+1),
				Skill:   rand.Float64() * 10000,
				Latency: rand.Float64() * 100,
			}
			playerNum++

			err := sendRequest(player)
			if err != nil {
				fmt.Printf("Failed to send request: %v\n", err)
			}
		}
		time.Sleep(time.Second)
	}

}
