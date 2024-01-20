package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gotrika/gotrika_backend/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type APIHandler struct {
	services *service.Services
}

func NewAPIHandler(services *service.Services) *APIHandler {
	return &APIHandler{
		services: services,
	}
}

func (h *APIHandler) Init(router *gin.Engine) {
	api := router.Group("/api")
	h.initUsersHandlers(api)
}

func parseIdFromPath(c *gin.Context, param string) (primitive.ObjectID, error) {
	idParam := c.Param(param)
	if idParam == "" {
		return primitive.ObjectID{}, errors.New("empty id param")
	}

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return primitive.ObjectID{}, errors.New("invalid id param")
	}

	return id, nil
}
