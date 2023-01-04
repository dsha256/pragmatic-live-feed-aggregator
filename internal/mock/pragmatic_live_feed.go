// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pragmaticlivefeed/pragmatic_live_feed.go

// Package mock_pragmaticlivefeed is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	sync "sync"

	dto "github.com/dsha256/pragmatic-live-feed-aggregator/pkg/dto"
	gomock "github.com/golang/mock/gomock"
	websocket "github.com/gorilla/websocket"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AddTable mock base method.
func (m *MockService) AddTable(ctx context.Context, table dto.PragmaticTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTable", ctx, table)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTable indicates an expected call of AddTable.
func (mr *MockServiceMockRecorder) AddTable(ctx, table interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTable", reflect.TypeOf((*MockService)(nil).AddTable), ctx, table)
}

// GetTableByTableAndCurrencyIDs mock base method.
func (m *MockService) GetTableByTableAndCurrencyIDs(ctx context.Context, tableID, currencyID string) (dto.PragmaticTable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTableByTableAndCurrencyIDs", ctx, tableID, currencyID)
	ret0, _ := ret[0].(dto.PragmaticTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTableByTableAndCurrencyIDs indicates an expected call of GetTableByTableAndCurrencyIDs.
func (mr *MockServiceMockRecorder) GetTableByTableAndCurrencyIDs(ctx, tableID, currencyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTableByTableAndCurrencyIDs", reflect.TypeOf((*MockService)(nil).GetTableByTableAndCurrencyIDs), ctx, tableID, currencyID)
}

// ListTables mock base method.
func (m *MockService) ListTables(ctx context.Context) ([]dto.PragmaticTableWithID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTables", ctx)
	ret0, _ := ret[0].([]dto.PragmaticTableWithID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTables indicates an expected call of ListTables.
func (mr *MockServiceMockRecorder) ListTables(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTables", reflect.TypeOf((*MockService)(nil).ListTables), ctx)
}

// MockPusherService is a mock of PusherService interface.
type MockPusherService struct {
	ctrl     *gomock.Controller
	recorder *MockPusherServiceMockRecorder
}

// MockPusherServiceMockRecorder is the mock recorder for MockPusherService.
type MockPusherServiceMockRecorder struct {
	mock *MockPusherService
}

// NewMockPusherService creates a new mock instance.
func NewMockPusherService(ctrl *gomock.Controller) *MockPusherService {
	mock := &MockPusherService{ctrl: ctrl}
	mock.recorder = &MockPusherServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPusherService) EXPECT() *MockPusherServiceMockRecorder {
	return m.recorder
}

// StartPushing mock base method.
func (m *MockPusherService) StartPushing(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartPushing", ctx)
}

// StartPushing indicates an expected call of StartPushing.
func (mr *MockPusherServiceMockRecorder) StartPushing(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartPushing", reflect.TypeOf((*MockPusherService)(nil).StartPushing), ctx)
}

// MockWSService is a mock of WSService interface.
type MockWSService struct {
	ctrl     *gomock.Controller
	recorder *MockWSServiceMockRecorder
}

// MockWSServiceMockRecorder is the mock recorder for MockWSService.
type MockWSServiceMockRecorder struct {
	mock *MockWSService
}

// NewMockWSService creates a new mock instance.
func NewMockWSService(ctrl *gomock.Controller) *MockWSService {
	mock := &MockWSService{ctrl: ctrl}
	mock.recorder = &MockWSServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWSService) EXPECT() *MockWSServiceMockRecorder {
	return m.recorder
}

// PushReceivedDataToDB mock base method.
func (m *MockWSService) PushReceivedDataToDB(conn *websocket.Conn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PushReceivedDataToDB", conn)
}

// PushReceivedDataToDB indicates an expected call of PushReceivedDataToDB.
func (mr *MockWSServiceMockRecorder) PushReceivedDataToDB(conn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushReceivedDataToDB", reflect.TypeOf((*MockWSService)(nil).PushReceivedDataToDB), conn)
}

// StartClient mock base method.
func (m *MockWSService) StartClient(msg []byte, wg *sync.WaitGroup) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartClient", msg, wg)
}

// StartClient indicates an expected call of StartClient.
func (mr *MockWSServiceMockRecorder) StartClient(msg, wg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartClient", reflect.TypeOf((*MockWSService)(nil).StartClient), msg, wg)
}

// StartClients mock base method.
func (m *MockWSService) StartClients() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartClients")
}

// StartClients indicates an expected call of StartClients.
func (mr *MockWSServiceMockRecorder) StartClients() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartClients", reflect.TypeOf((*MockWSService)(nil).StartClients))
}