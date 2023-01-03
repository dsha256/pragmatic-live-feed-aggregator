package repository

import (
	"context"
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/dto"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/utils"
)

// TODO: Improve error-handling as will use logger

type MemcachedLiveFeedRepository struct {
	client *memcache.Client
	keys   []string
}

func NewMemcachedLiveFeedRepository(client *memcache.Client) *MemcachedLiveFeedRepository {
	return &MemcachedLiveFeedRepository{client: client}
}

func (db *MemcachedLiveFeedRepository) AddTable(ctx context.Context, table dto.PragmaticTable) error {
	itemToAddKey := utils.GenerateIDFromTableAndCurrencyIDs(table.TableId, table.Currency)
	marshaledTable, err := json.Marshal(table)
	if err != nil {
		return err
	}
	itemToAdd := &memcache.Item{
		Key:   itemToAddKey,
		Value: marshaledTable,
	}
	err = db.client.Set(itemToAdd)
	if err != nil {
		return err
	}
	db.keys = append(db.keys, itemToAddKey)
	return nil
}

func (db *MemcachedLiveFeedRepository) GetTableByTableAndCurrencyIDs(ctx context.Context, tableID, currencyID string) (dto.PragmaticTable, error) {
	itemToGetID := utils.GenerateIDFromTableAndCurrencyIDs(tableID, currencyID)
	table, err := db.client.Get(itemToGetID)
	if err != nil {
		return dto.PragmaticTable{}, err
	}
	var pragmaticTable dto.PragmaticTable
	err = json.Unmarshal(table.Value, &pragmaticTable)
	return pragmaticTable, nil
}

func (db *MemcachedLiveFeedRepository) ListTables(ctx context.Context) ([]dto.PragmaticTableWithID, error) {
	var pragmaticTables []dto.PragmaticTableWithID
	_ = pragmaticTables

	for _, key := range db.keys {
		var pragmaticTable dto.PragmaticTable
		item, err := db.client.Get(key)
		if err != nil {
			return []dto.PragmaticTableWithID{}, err
		}
		err = json.Unmarshal(item.Value, &pragmaticTables)
		if err != nil {
			return []dto.PragmaticTableWithID{}, err
		}
		pragmaticTables = append(pragmaticTables, dto.PragmaticTableWithID{
			TableAndCurrencyID: key,
			PragmaticTable:     pragmaticTable,
		})
	}

	return pragmaticTables, nil
}
