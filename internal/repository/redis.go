package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/dto"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/utils"
	"github.com/go-redis/redis/v8"
)

var (
	ErrKeyDoesNotExist = errors.New("key does not exist")
	ErrRedisInternal   = errors.New("redis internal error")
)

type RedisLiveFeedRepository struct {
	client *redis.Client
}

func NewRedisLiveFeedRepository(client *redis.Client) *RedisLiveFeedRepository {
	return &RedisLiveFeedRepository{
		client,
	}
}

func (db *RedisLiveFeedRepository) AddTable(ctx context.Context, table dto.PragmaticTable) error {
	redisID := utils.GenerateIDFromTableAndCurrencyIDs(table.TableId, table.Currency)

	jsonPragmaticTable, err := json.Marshal(table)
	if err != nil {
		return err
	}

	db.client.Set(ctx, redisID, jsonPragmaticTable, 0)
	return nil
}

func (db *RedisLiveFeedRepository) GetTableByTableAndCurrencyIDs(ctx context.Context, tableID, currencyID string) (dto.PragmaticTable, error) {
	var pragmaticTable dto.PragmaticTable

	tableUniqueID := utils.GenerateIDFromTableAndCurrencyIDs(tableID, currencyID)
	table, err := db.client.Get(ctx, tableUniqueID).Result()
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

func (db *RedisLiveFeedRepository) ListTables(ctx context.Context) ([]dto.PragmaticTableWithID, error) {
	var pragmaticTables []dto.PragmaticTableWithID
	var cursor uint64
	var table string

	keys, cursor, err := db.client.Scan(ctx, cursor, "*:*", 0).Result()
	if err != nil {
		return pragmaticTables, err
	}

	for _, key := range keys {
		table, err = db.client.Get(ctx, key).Result()
		if err != nil {
			return []dto.PragmaticTableWithID{}, err
		}
		var pragmaticTableWithID dto.PragmaticTableWithID
		var pragmaticTable dto.PragmaticTable
		err := json.Unmarshal([]byte(table), &pragmaticTable)
		if err != nil {
			return []dto.PragmaticTableWithID{}, err
		}
		pragmaticTableWithID.TableAndCurrencyID = key
		pragmaticTableWithID.PragmaticTable = pragmaticTable
		pragmaticTables = append(pragmaticTables, pragmaticTableWithID)
	}

	return pragmaticTables, nil
}
