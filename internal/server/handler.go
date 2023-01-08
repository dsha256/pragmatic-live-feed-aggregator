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

// PragmaticTable lists Pragmatic live feed tables data.
// @Summary List Pragmatic live feed tables data
// @Schemes
// @Description List Pragmatic live feed tables data
// @Accept json
// @Produce json
// @Success 200 {object} dto.Response{data=[]dto.PragmaticTableWithID}
// @Router /tables [get]
func (h *pragmaticLiveFeedHandler) PragmaticTable(c *gin.Context) {
	pragmaticTables, err := h.service.ListTables(c)
	if err != nil {
		handleErrResp(c, "error getting pragmatic live feed tables data from the db", http.StatusInternalServerError)
		return
	}
	handleSuccessfulResp(c, "", pragmaticTables)
}

// Health checks if the service is up
// @Summary Check if the service is up
// @Schemes
// @Description Check if the service is up
// @Accept json
// @Produce json
// @Success 200 {object} dto.Response
// @Router /tables/health [get]
func (h *pragmaticLiveFeedHandler) Health(c *gin.Context) {
	handleSuccessfulResp(c, "working...", nil)
}
