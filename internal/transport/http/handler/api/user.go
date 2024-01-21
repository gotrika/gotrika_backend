package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gotrika/gotrika_backend/internal/dto"
)

func (h *APIHandler) initUsersHandlers(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/sign-up/", h.userSignUp)
		auth.POST("/sign-in/", h.userSignIn)
		auth.POST("/refresh/", h.userUpdateTokens)
	}
	users := api.Group("users", authMiddleware(h.services.Users))
	{
		users.GET("/me/", h.userMe)
		users.GET("/:user_id/", h.userGet)
	}

}

func (h *APIHandler) userSignUp(c *gin.Context) {
	var inp dto.RegisterUserDTO
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	userID, err := h.services.Users.SignUp(c.Request.Context(), inp)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	resp := idResponse{ID: userID}
	c.JSON(http.StatusCreated, &resp)
}

func (h *APIHandler) userSignIn(c *gin.Context) {
	var inp dto.AuthUserDTO
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	authResponse, err := h.services.Users.SignIn(c.Request.Context(), inp)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &authResponse)
}

func (h *APIHandler) userUpdateTokens(c *gin.Context) {
	refreshToken := c.DefaultQuery("token", "")
	if refreshToken == "" {
		newResponse(c, http.StatusForbidden, "miss refresh token")
		return
	}
	authResponse, err := h.services.Users.UpdateTokens(c.Request.Context(), refreshToken)
	if err != nil {
		newResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, &authResponse)
}

func (h *APIHandler) userGet(c *gin.Context) {
	userID, err := parseIdFromPath(c, "user_id")
	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid user_id")
		return
	}
	user, _, err := h.services.Users.GetUserByID(c.Request.Context(), userID.Hex())
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &user)
}

func (h *APIHandler) userMe(c *gin.Context) {
	userID := c.GetString(userIDCtx)
	user, _, err := h.services.Users.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &user)
}
