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
	err := s.repo.AddTable(ctx, table)
	if err != nil {
		return err
	}
	return nil
}

func (s service) GetTableByTableAndCurrencyIDs(ctx context.Context, tableID, currencyID string) (dto.PragmaticTable, error) {
	pt, err := s.repo.GetTableByTableAndCurrencyIDs(ctx, tableID, currencyID)
	if err != nil {
		return dto.PragmaticTable{}, err
	}
	return pt, nil
}

func (s service) ListTables(ctx context.Context) ([]dto.PragmaticTableWithID, error) {
	// TODO: Refactor to move extra processing from db layer to here to avoid locking db for a long time
	pts, err := s.repo.ListTables(ctx)
	if err != nil {
		return nil, err
	}
	return pts, nil
}
