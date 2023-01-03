package pragmaticlivefeed

import (
	"context"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/dto"
	"github.com/gorilla/websocket"
	"sync"
)

type Service interface {
	AddTable(ctx context.Context, table dto.PragmaticTable) error
	GetTableByTableAndCurrencyIDs(ctx context.Context, tableID, currencyID string) (dto.PragmaticTable, error)
	ListTables(ctx context.Context) ([]dto.PragmaticTableWithID, error)
}

type PusherService interface {
	StartPushing(ctx context.Context)
}

type WSService interface {
	StartClients()
	StartClient(msg []byte, wg *sync.WaitGroup)
	PushReceivedDataToDB(conn *websocket.Conn)
}
