package server

import (
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/repo"
	"github.com/gin-gonic/gin"
)

type pragmaticTableHandler struct {
	repo *repo.InMemoryDataBase
}

func newPragmaticTableHandler(repo repo.InMemoryDataBase) *pragmaticTableHandler {
	return &pragmaticTableHandler{repo: &repo}
}

func (h *pragmaticTableHandler) registerRoutes(r *gin.RouterGroup) {
	r.GET("/tables/health", h.Health)
	r.GET("/tables", h.PragmaticTable)
}

func (h *pragmaticTableHandler) PragmaticTable(c *gin.Context) {
	// TODO: Implement an appropriate method om the DB layer
}

func (h *pragmaticTableHandler) Health(c *gin.Context) {
	HandleSuccessfulResp(c, "working...", nil)
}
