package server

import (
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/pragmaticlivefeed"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTP struct {
	http.Handler
	service pragmaticlivefeed.Service
	engine  *gin.Engine
}

func NewHTTP(pragmaticLiveFeedSvc pragmaticlivefeed.Service) *HTTP {
	engine := gin.New()
	// TODO: fix as will have do mains
	engine.Use(CORS())
	server := &HTTP{Handler: engine, engine: engine}

	pragmaticLiveFeedRoute := engine.Group("api/v1/pragmatic_live_feed")

	pragmaticTablesHandler := newPragmaticTableHandler(pragmaticLiveFeedSvc)
	pragmaticTablesHandler.registerRoutes(pragmaticLiveFeedRoute)

	return server
}
