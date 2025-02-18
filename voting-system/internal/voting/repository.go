package voting

import (
	"context"
	"errors"
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
	"time"
	"voting-system/internal/redis"
)

var ctx = context.Background()

func CreatePoll(pollID string, options []string, expiration time.Duration) error {
	for _, option := range options {
		key := fmt.Sprintf("poll:%s:%s", pollID, option)
		_, err := redis.RedisClient.Set(ctx, key, 0, 0).Result()
		if err != nil {
			return fmt.Errorf("could not create poll option %s: %v", option, err)
		}
	}

	pollKey := fmt.Sprintf("poll:%s:expiration", pollID)
	_, err := redis.RedisClient.Set(ctx, pollKey, time.Now().Add(expiration).Unix(), 0).Result()
	if err != nil {
		return fmt.Errorf("could not set expiration for poll %s: %v", pollID, err)
	}
	return nil
}

func AddVote(pollID, option string) error {
	key := fmt.Sprintf("poll:%s:%s", pollID, option)
	_, err := redis.RedisClient.Incr(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("could not increment vote for option %s: %v", option, err)
	}
	return nil
}

func GetPollResults(pollID string) (map[string]int64, error) {
	options := []string{"A", "B", "C"}
	results := make(map[string]int64)

	for _, option := range options {
		key := fmt.Sprintf("poll:%s:%s", pollID, option)
		count, err := redis.RedisClient.Get(ctx, key).Int64()
		if err != nil {
			if !errors.Is(err, redis2.Nil) {
				return nil, fmt.Errorf("could not fetch vote count for option %s: %v", option, err)
			}
			count = 0
		}
		results[option] = count
	}
	return results, nil
}
