package api

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gotrika/gotrika_backend/internal/service"
	"github.com/gotrika/gotrika_backend/internal/transport/http/handler/utils"
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
	h.initSitesHandlers(api)
}

func parseIdFromPath(c *gin.Context, param string) (primitive.ObjectID, error) {
	idParam := c.Param(param)
	if idParam == "" {
		return primitive.ObjectID{}, errors.New("empty id param")
	}

	return utils.ConverIDtoObjectId(idParam)
}

func getLimitOffsetFromQueryParams(c *gin.Context) (int, int) {
	limitParam := c.DefaultQuery("limit", "100")
	offsetParam := c.DefaultQuery("offset", "0")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return 100, 0
	}
	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		return 100, 0
	}
	return limit, offset

}
