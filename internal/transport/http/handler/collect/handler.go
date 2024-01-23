package collect

import (
	"github.com/gin-gonic/gin"
	"github.com/gotrika/gotrika_backend/internal/service"
)

type CollectHandler struct {
	services *service.Services
}

func NewCollectHandler(services *service.Services) *CollectHandler {
	return &CollectHandler{
		services: services,
	}
}

func (h *CollectHandler) Init(router *gin.Engine) {
	collect := router.Group("/collect")
	h.initTrackerHandlers(collect)
}
