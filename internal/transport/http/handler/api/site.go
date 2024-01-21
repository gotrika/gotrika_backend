package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gotrika/gotrika_backend/internal/core"
	"github.com/gotrika/gotrika_backend/internal/dto"
)

func (h *APIHandler) initSitesHandlers(api *gin.RouterGroup) {
	sites := api.Group("/sites", authMiddleware(h.services.Users))
	{
		sites.POST("/", h.sitesCreate)
		sites.GET("/", h.siteList)
		sites.GET("/:site_id/", h.sitesGet)
		sites.PUT("/:site_id/", h.siteUpdate)
		sites.DELETE("/:site_id/", h.siteDelete)
	}
}

func (h *APIHandler) sitesCreate(c *gin.Context) {
	var inp dto.CreateSiteDTO
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input site body")
		return
	}
	userID := c.GetString(userIDCtx)
	objID, _ := converIDtoObjectId(userID)
	resp, err := h.services.Sites.CreateSite(c.Request.Context(), objID, inp)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &resp)
}

func (h *APIHandler) sitesGet(c *gin.Context) {
	siteID, err := parseIdFromPath(c, "site_id")
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userID := c.GetString(userIDCtx)
	objID, _ := converIDtoObjectId(userID)
	isAdmin := c.GetBool(userIsAdminCtx)
	siteRetrieve, err := h.services.Sites.GetSiteByID(c.Request.Context(), isAdmin, objID, siteID)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &siteRetrieve)
}

func (h *APIHandler) siteDelete(c *gin.Context) {
	siteID, err := parseIdFromPath(c, "site_id")
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userID := c.GetString(userIDCtx)
	objID, _ := converIDtoObjectId(userID)
	isAdmin := c.GetBool(userIsAdminCtx)
	err = h.services.Sites.DeleteSite(c.Request.Context(), isAdmin, objID, siteID)
	if err != nil {
		newResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (s *APIHandler) siteList(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	search := queryParams.Get("search")
	limit, offset := getLimitOffsetFromQueryParams(c)
	userID := c.GetString(userIDCtx)
	objID, _ := converIDtoObjectId(userID)
	isAdmin := c.GetBool(userIsAdminCtx)
	data, count, err := s.services.Sites.ListSites(
		c.Request.Context(),
		search,
		isAdmin,
		objID,
		limit,
		offset,
	)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, dataResponse{Data: data, Count: count})
}

func (s *APIHandler) siteUpdate(c *gin.Context) {
	siteID, err := parseIdFromPath(c, "site_id")
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var siteDTO dto.UpdateSiteDTO
	if err := c.BindJSON(&siteDTO); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input site body")
		return
	}
	userID := c.GetString(userIDCtx)
	objID, _ := converIDtoObjectId(userID)
	isAdmin := c.GetBool(userIsAdminCtx)
	res, err := s.services.Sites.UpdateSite(c.Request.Context(), isAdmin, objID, siteID, &siteDTO)
	if err != nil {
		newResponse(c, http.StatusForbidden, core.ErrSiteAccessDenied.Error())
		return
	}
	c.JSON(http.StatusOK, res)

}
