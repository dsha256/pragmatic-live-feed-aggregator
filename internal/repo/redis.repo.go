package repo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/dto"
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	ErrKeyDoesNotExist = errors.New("key does not exist")
	ErrRedisInternal   = errors.New("redis internal error")
)

type RedisRepository struct {
	sync.RWMutex
	Client *redis.Client
}

func NewRedisRepository(client *redis.Client) RedisRepository {
	return RedisRepository{
		sync.RWMutex{},
		client,
	}
}

func (db *RedisRepository) AddTable(ctx context.Context, table dto.PragmaticTable) error {
	db.Lock()
	defer db.Unlock()

	redisID := table.CreateTableCurrencyID()

	jsonPragmaticTable, err := json.Marshal(table)
	if err != nil {
		return err
	}

	db.Client.Set(ctx, redisID, jsonPragmaticTable, 0)
	return nil
}

func (db *RedisRepository) GetTableByTableAndCurrencyIDs(ctx context.Context, tableID, currencyID string) (dto.PragmaticTable, error) {
	db.Lock()
	defer db.Unlock()

	var pragmaticTable dto.PragmaticTable

	tableUniqueID := tableID + currencyID
	table, err := db.Client.Get(ctx, tableUniqueID).Result()
	switch {
	case err == redis.Nil:
		return dto.PragmaticTable{}, ErrKeyDoesNotExist
	case err != nil:
		return dto.PragmaticTable{}, ErrRedisInternal
	}

	err = json.Unmarshal([]byte(table), &pragmaticTable)
	if err != nil {
		return dto.PragmaticTable{}, errors.New("failed to unmarshal the data retrieved from Redis: " + err.Error())
	}

	return pragmaticTable, nil
}

func (db *RedisRepository) ListTables(ctx context.Context) ([]dto.PragmaticTable, error) {
	db.Lock()
	defer db.Unlock()

	var pragmaticTables []dto.PragmaticTable
	// TODO: Implement me

	return pragmaticTables, nil
}
