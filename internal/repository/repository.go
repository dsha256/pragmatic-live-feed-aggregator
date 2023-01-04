package repository

import (
	"github.com/go-redis/redis/v8"
)

type Aggregator struct {
	*MapLiveFeedRepository
	db *redis.Client
}

func NewAggregator() *Aggregator {
	return &Aggregator{MapLiveFeedRepository: NewMapLiveFeedRepository()}
}
