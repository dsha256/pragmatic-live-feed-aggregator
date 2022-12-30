package server

import (
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/repo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type pragmaticTableHandler struct {
	repo *repo.RedisRepository
}

func newPragmaticTableHandler(repo *repo.RedisRepository) *pragmaticTableHandler {
	return &pragmaticTableHandler{repo: repo}
}

func (h *pragmaticTableHandler) registerRoutes(r *gin.RouterGroup) {
	r.GET("/tables/health", h.Health)
	r.GET("/tables", h.PragmaticTable)
}

func (h *pragmaticTableHandler) PragmaticTable(c *gin.Context) {
	// TODO: Implement an appropriate method om the DB layer
	pragmaticTables, err := h.repo.ListTables(c)
	if err != nil {
		handleErrResp(c, "error getting pragmatic tables data from the db", http.StatusInternalServerError)
		return
	}
	handleSuccessfulResp(c, "", pragmaticTables)
}

func (h *pragmaticTableHandler) Health(c *gin.Context) {
	handleSuccessfulResp(c, "working...", nil)
}
