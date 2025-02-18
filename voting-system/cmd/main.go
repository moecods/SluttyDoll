package main

import (
	"fmt"
	"time"
	"voting-system/internal/redis"
	"voting-system/internal/voting"
)

func main() {
	redis.InitRedis()

	pollID := "poll1"
	options := []string{"A", "B", "C"}
	expiration := 24 * time.Hour

	err := voting.CreatePoll(pollID, options, expiration)
	if err != nil {
		return
	}

	err = voting.AddVote(pollID, "A")
	if err != nil {
		return
	}

	err = voting.AddVote(pollID, "A")
	if err != nil {
		return
	}

	err = voting.AddVote(pollID, "B")
	if err != nil {
		return
	}

	results, err := voting.GetPollResults(pollID)
	if err != nil {
		return
	}

	for option, value := range results {
		fmt.Println(value, option)
	}
}
