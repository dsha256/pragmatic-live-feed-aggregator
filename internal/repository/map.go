package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/dto"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/utils"
	"sync"
)

type MapLiveFeedRepository struct {
	mutex  sync.RWMutex
	tables map[string]dto.PragmaticTable
}

func NewMapLiveFeedRepository() *MapLiveFeedRepository {
	return &MapLiveFeedRepository{
		tables: make(map[string]dto.PragmaticTable),
	}
}

func (m *MapLiveFeedRepository) AddTable(ctx context.Context, table dto.PragmaticTable) error {
	id := utils.GenerateIDFromTableAndCurrencyIDs(table.TableId, table.Currency)

	m.mutex.Lock()
	m.tables[id] = table
	m.mutex.Unlock()

	return nil
}

func (m *MapLiveFeedRepository) GetTableByTableAndCurrencyIDs(ctx context.Context, tableID, currencyID string) (dto.PragmaticTable, error) {
	uniqID := utils.GenerateIDFromTableAndCurrencyIDs(tableID, currencyID)

	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if table, ok := m.tables[uniqID]; ok {
		return table, nil
	}

	return dto.PragmaticTable{}, errors.New(fmt.Sprintf("No such key: %s", uniqID))
}

func (m *MapLiveFeedRepository) ListTables(ctx context.Context) ([]dto.PragmaticTableWithID, error) {
	var pragmaticTables []dto.PragmaticTableWithID

	m.mutex.RLock()
	for k, v := range m.tables {
		pragmaticTables = append(pragmaticTables, dto.PragmaticTableWithID{TableAndCurrencyID: k, PragmaticTable: v})
	}
	m.mutex.RUnlock()

	return pragmaticTables, nil
}
