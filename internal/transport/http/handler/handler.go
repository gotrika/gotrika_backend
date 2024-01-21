package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gotrika/gotrika_backend/docs"
	"github.com/gotrika/gotrika_backend/internal/config"
	"github.com/gotrika/gotrika_backend/internal/service"
	"github.com/gotrika/gotrika_backend/internal/transport/http/handler/api"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	// Init gin handler
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	apiHandler := api.NewAPIHandler(h.services)
	apiHandler.Init(router)
	return router
}
