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

// @Summary User Sign Up
// @Tags auth
// @Description create user account
// @ModuleID userSignUp
// @Accept  json
// @Produce  json
// @Param input body dto.RegisterUserDTO true "sign up info"
// @Success 201 {object} idResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/auth/sign-up/ [post]
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

// @Summary User Sign In
// @Tags auth
// @Description user sign in
// @ModuleID userSignIn
// @Accept  json
// @Produce  json
// @Param input body dto.AuthUserDTO true "sign up info"
// @Success 200 {object} dto.AuthResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/auth/sign-in/ [post]
func (h *APIHandler) userSignIn(c *gin.Context) {
	var inp dto.AuthUserDTO
	if err := c.Bind(&inp); err != nil {
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

// @Summary User Updates Tokens
// @Tags auth
// @Description user updates tokens
// @Accept  json
// @Produce  json
// @Param input body dto.UpdateTokensByRefreshToken true "sign up info"
// @Success 200 {object} dto.AuthResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/auth/refresh/ [post]
func (h *APIHandler) userUpdateTokens(c *gin.Context) {
	var inp dto.UpdateTokensByRefreshToken
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if inp.RefreshToken == "" {
		newResponse(c, http.StatusBadRequest, "empty refresh_token")
		return
	}
	authResponse, err := h.services.Users.UpdateTokens(c.Request.Context(), inp.RefreshToken)
	if err != nil {
		newResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, &authResponse)
}

// @Summary User Info
// @Security ApiAuth
// @Tags users
// @Description user info
// @Accept  json
// @Produce  json
// @Param user_id path string true "user id"
// @Success 200 {object} dto.UserRetrieveDTO
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/users/{user_id}/ [get]
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

// @Summary Current User Info
// @Security ApiAuth
// @Tags users
// @Description current user info
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.UserRetrieveDTO
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/users/me/ [get]
func (h *APIHandler) userMe(c *gin.Context) {
	userID := c.GetString(userIDCtx)
	user, _, err := h.services.Users.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &user)
}
