package pragmaticlivefeed

import (
	"context"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/dto"
)

type service struct {
	repo AggregateRepository
}

func NewService(repo AggregateRepository) *service {
	return &service{repo: repo}
}

func (s service) AddTable(ctx context.Context, table dto.PragmaticTable) error {
	return s.repo.AddTable(ctx, table)
}

func (s service) GetTableByTableAndCurrencyIDs(ctx context.Context, tableID, currencyID string) (dto.PragmaticTable, error) {
	return s.repo.GetTableByTableAndCurrencyIDs(ctx, tableID, currencyID)
}

func (s service) ListTables(ctx context.Context) ([]dto.PragmaticTableWithID, error) {
	return s.repo.ListTables(ctx)
}
