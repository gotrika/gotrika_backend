package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gotrika/gotrika_backend/pkg/logger"
)

type dataResponse struct {
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

type idResponse struct {
	ID interface{} `json:"id"`
}

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
