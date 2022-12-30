package server

import (
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/repo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTP struct {
	http.Handler
	repo   repo.InMemoryDataBase
	engine *gin.Engine
}

func NewHTTP(
	repo *repo.RedisRepository,
) *HTTP {
	engine := gin.New()
	// TODO: fix as will have domains
	engine.Use(CORS())
	server := &HTTP{
		Handler: engine,
		engine:  engine,
	}
	pragmaticNewsFeedRoute := engine.Group("pragmatic_news_feed/v1")

	pragmaticTablesHandler := newPragmaticTableHandler(repo)
	pragmaticTablesHandler.registerRoutes(pragmaticNewsFeedRoute)

	return server
}
