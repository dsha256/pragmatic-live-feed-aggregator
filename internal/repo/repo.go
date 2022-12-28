package repo

import (
	"context"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/dto"
)

type InMemoryDataBase interface {
	AddTable(ctx context.Context, table dto.PragmaticTable) error
	GetTableByTableAndCurrencyIDs(ctx context.Context, tableID, currencyID string) (dto.PragmaticTable, error)
	ListTables(ctx context.Context) ([]dto.PragmaticTable, error)
}
