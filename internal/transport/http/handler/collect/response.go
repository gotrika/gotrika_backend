package collect

import "github.com/gin-gonic/gin"

type trackerResponse struct {
	Success bool `json:"success"`
}

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, response{message})
}
