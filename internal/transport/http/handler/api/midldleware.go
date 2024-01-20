package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gotrika/gotrika_backend/internal/service"
)

const (
	authorizationHeader = "Authorization"
	userIDCtx           = "userID"
)

func extractAuthToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("miss auth header")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}
	if strings.ToLower(jwtToken[0]) != "jwt" {
		return "", errors.New("incorrectly formatted authorization header param")
	}
	return jwtToken[1], nil
}

func authMiddleware(userService service.Users) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken, err := extractAuthToken(c.GetHeader(authorizationHeader))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response{
				Message: err.Error(),
			})
			return
		}
		userID, sign, err := userService.TokenManager().Parse(authToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response{
				Message: err.Error(),
			})
			return
		}
		_, userSign, err := userService.GetUserByID(c, userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response{
				Message: err.Error(),
			})
			return
		}
		if userSign != sign {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response{
				Message: "invalid sign",
			})
			return
		}
		c.Set(userIDCtx, userID)
		c.Next()

	}
}
