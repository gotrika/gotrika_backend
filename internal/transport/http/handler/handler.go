package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gotrika/gotrika_backend/internal/config"
	"github.com/gotrika/gotrika_backend/internal/service"
	"github.com/gotrika/gotrika_backend/internal/transport/http/handler/api"
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

	router.GET("/ping/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	apiHandler := api.NewAPIHandler(h.services)
	apiHandler.Init(router)
	return router
}
