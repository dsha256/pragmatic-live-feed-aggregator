package server

import (
	"encoding/json"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/mock"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/dto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPragmaticLiveFeedHandler_Health(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name          string
		method        string
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK when service is up and running",
			method: http.MethodGet,
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

				expectedResp := dto.Response{
					Data:    nil,
					Error:   false,
					Message: "working...",
					Code:    200,
					Status:  200,
				}
				actualRespBt := recorder.Body.Bytes()
				var actualResp dto.Response
				err := json.Unmarshal(actualRespBt, &actualResp)
				require.NoError(t, err)
				require.Equal(t, actualResp, expectedResp)
			},
		},
		{
			name:   "404 on HTTP request POST",
			method: http.MethodPost,
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "404 on HTTP request PUT",
			method: http.MethodPut,
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "404 on HTTP request PATCH",
			method: http.MethodPatch,
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "404 on HTTP request DELETE",
			method: http.MethodDelete,
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "404 on HTTP request CONNECT",
			method: http.MethodConnect,
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "404 on HTTP request TRACE",
			method: http.MethodTrace,
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "404 on HTTP request OPTIONS",
			method: http.MethodOptions,
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mock.NewMockService(ctrl)

			server := NewHTTP(service)
			recorder := httptest.NewRecorder()

			url := "/api/v1/pragmatic_live_feed/tables/health"
			req, err := http.NewRequest(tc.method, url, nil)
			require.NoError(t, err)

			server.Handler.ServeHTTP(recorder, req)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestPragmaticLiveFeedHandler_PragmaticTable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name          string
		method        string
		buildStubs    func(service *mock.MockService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK on HTTP request POST",
			method: http.MethodGet,
			buildStubs: func(service *mock.MockService) {
				service.
					EXPECT().
					ListTables(gomock.Any()).Times(1).
					Return([]dto.PragmaticTableWithID{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:   "404 on HTTP request POST",
			method: http.MethodPost,
			buildStubs: func(service *mock.MockService) {
				service.
					EXPECT().
					ListTables(gomock.Any()).Times(0).
					Return([]dto.PragmaticTableWithID{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "404 on HTTP request PUT",
			method: http.MethodPut,
			buildStubs: func(service *mock.MockService) {
				service.
					EXPECT().
					ListTables(gomock.Any()).Times(0).
					Return([]dto.PragmaticTableWithID{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "404 on HTTP request PATCH",
			method: http.MethodPatch,
			buildStubs: func(service *mock.MockService) {
				service.
					EXPECT().
					ListTables(gomock.Any()).Times(0).
					Return([]dto.PragmaticTableWithID{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "404 on HTTP request DELETE",
			method: http.MethodDelete,
			buildStubs: func(service *mock.MockService) {
				service.
					EXPECT().
					ListTables(gomock.Any()).Times(0).
					Return([]dto.PragmaticTableWithID{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "404 on HTTP request CONNECT",
			method: http.MethodConnect,
			buildStubs: func(service *mock.MockService) {
				service.
					EXPECT().
					ListTables(gomock.Any()).Times(0).
					Return([]dto.PragmaticTableWithID{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "404 on HTTP request TRACE",
			method: http.MethodTrace,
			buildStubs: func(service *mock.MockService) {
				service.
					EXPECT().
					ListTables(gomock.Any()).Times(0).
					Return([]dto.PragmaticTableWithID{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "404 on HTTP request OPTIONS",
			method: http.MethodOptions,
			buildStubs: func(service *mock.MockService) {
				service.
					EXPECT().
					ListTables(gomock.Any()).Times(0).
					Return([]dto.PragmaticTableWithID{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mock.NewMockService(ctrl)
			tc.buildStubs(service)

			server := NewHTTP(service)
			recorder := httptest.NewRecorder()

			url := "/api/v1/pragmatic_live_feed/tables"
			req, err := http.NewRequest(tc.method, url, nil)
			require.NoError(t, err)

			server.Handler.ServeHTTP(recorder, req)
			tc.checkResponse(t, recorder)
		})
	}
}
