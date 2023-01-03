package repository

import (
	"github.com/go-redis/redis/v8"
)

type Aggregator struct {
	*RedisLiveFeedRepository
	db *redis.Client
}

func NewAggregator(db *redis.Client) *Aggregator {
	return &Aggregator{RedisLiveFeedRepository: NewRedisLiveFeedRepository(db), db: db}
}

// InTx runs passed func in tx.
//func (a *Aggregator) InTx(ctx context.Context, f pragmaticlivefeed.TxF) error {
//}
