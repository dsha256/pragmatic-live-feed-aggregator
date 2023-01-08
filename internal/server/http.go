package server

import (
	_ "github.com/dsha256/pragmatic-live-feed-aggregator/docs/swagger"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/pragmaticlivefeed"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type HTTP struct {
	http.Handler
	service pragmaticlivefeed.Service
	engine  *gin.Engine
}

func NewHTTP(pragmaticLiveFeedSvc pragmaticlivefeed.Service) *HTTP {
	engine := gin.New()
	// TODO: fix as will have domains
	engine.Use(CORS())
	server := &HTTP{Handler: engine, engine: engine}

	// Swagger UI
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	pragmaticLiveFeedRoute := engine.Group("api/v1/pragmatic_live_feed")

	pragmaticTablesHandler := newPragmaticTableHandler(pragmaticLiveFeedSvc)
	pragmaticTablesHandler.registerRoutes(pragmaticLiveFeedRoute)

	return server
}
