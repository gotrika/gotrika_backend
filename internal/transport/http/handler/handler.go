package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gotrika/gotrika_backend/docs"
	"github.com/gotrika/gotrika_backend/internal/config"
	"github.com/gotrika/gotrika_backend/internal/service"
	"github.com/gotrika/gotrika_backend/internal/transport/http/handler/api"
	"github.com/mvrilo/go-redoc"
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
	rdoc := redoc.Redoc{
		Title:       "Gotrika API",
		Description: "Gotrika Backend API",
		SpecFile:    "./docs/swagger.json",
		SpecPath:    "/docs/doc.json",
		DocsPath:    "/redoc/",
	}

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.GET("/redoc/", func(ctx *gin.Context) {
		handler := rdoc.Handler()
		handler(ctx.Writer, ctx.Request)
		ctx.Next()
	})
	apiHandler := api.NewAPIHandler(h.services)
	apiHandler.Init(router)
	return router
}
