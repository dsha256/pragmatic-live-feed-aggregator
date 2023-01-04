package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/dto"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/utils"
)

type MapLiveFeedRepository struct {
	tables map[string]dto.PragmaticTable
}

func (m *MapLiveFeedRepository) AddTable(ctx context.Context, table dto.PragmaticTable) error {
	id := utils.GenerateIDFromTableAndCurrencyIDs(table.TableId, table.Currency)
	m.tables[id] = table
	return nil
}

func (m *MapLiveFeedRepository) GetTableByTableAndCurrencyIDs(ctx context.Context, tableID, currencyID string) (dto.PragmaticTable, error) {
	uniqID := utils.GenerateIDFromTableAndCurrencyIDs(tableID, currencyID)
	if table, ok := m.tables[uniqID]; ok {
		return table, nil
	}
	return dto.PragmaticTable{}, errors.New(fmt.Sprintf("No such key: %s", uniqID))
}

func (m *MapLiveFeedRepository) ListTables(ctx context.Context) ([]dto.PragmaticTableWithID, error) {
	var pragmaticTables []dto.PragmaticTableWithID
	for k, v := range m.tables {
		pragmaticTables = append(pragmaticTables, dto.PragmaticTableWithID{TableAndCurrencyID: k, PragmaticTable: v})
	}
	return pragmaticTables, nil
}
