package server

import (
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/pragmaticlivefeed"
	"github.com/gin-gonic/gin"
	"net/http"
)

type pragmaticLiveFeedHandler struct {
	service pragmaticlivefeed.Service
}

func newPragmaticTableHandler(service pragmaticlivefeed.Service) *pragmaticLiveFeedHandler {
	return &pragmaticLiveFeedHandler{service: service}
}

func (h *pragmaticLiveFeedHandler) registerRoutes(r *gin.RouterGroup) {
	r.GET("/tables/health", h.Health)
	r.GET("/tables", h.PragmaticTable)
}

func (h *pragmaticLiveFeedHandler) PragmaticTable(c *gin.Context) {
	pragmaticTables, err := h.service.ListTables(c)
	if err != nil {
		handleErrResp(c, "error getting pragmatic live feed tables data from the db", http.StatusInternalServerError)
		return
	}
	handleSuccessfulResp(c, "", pragmaticTables)
}

func (h *pragmaticLiveFeedHandler) Health(c *gin.Context) {
	handleSuccessfulResp(c, "working...", nil)
}
