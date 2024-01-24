package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gotrika/gotrika_backend/internal/core"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"github.com/gotrika/gotrika_backend/internal/transport/http/handler/utils"
)

type sitesResponse struct {
	Count int                    `json:"count"`
	Data  []*dto.SiteRetrieveDTO `json:"data"`
}

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

// @Summary Add Site
// @Tags sites
// @Description add site
// @ModuleID sites
// @Security ApiAuth
// @Accept  json
// @Produce  json
// @Param input body dto.CreateSiteDTO true "site info"
// @Success 200 {object} dto.SiteRetrieveDTO
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/sites/ [post]
func (h *APIHandler) sitesCreate(c *gin.Context) {
	var inp dto.CreateSiteDTO
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input site body")
		return
	}
	url := fmt.Sprintf("http://%s", c.Request.Host)
	userID := c.GetString(userIDCtx)
	objID, _ := utils.ConverIDtoObjectId(userID)
	resp, err := h.services.Sites.CreateSite(c.Request.Context(), url, objID, inp)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &resp)
}

// @Summary Site Info
// @Security ApiAuth
// @Tags sites
// @Description site info
// @Accept  json
// @Produce  json
// @Param site_id path string true "site id"
// @Success 200 {object} dto.SiteRetrieveDTO
// @Failure 400,403 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/sites/{site_id}/ [get]
func (h *APIHandler) sitesGet(c *gin.Context) {
	siteID, err := parseIdFromPath(c, "site_id")
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userID := c.GetString(userIDCtx)
	objID, _ := utils.ConverIDtoObjectId(userID)
	isAdmin := c.GetBool(userIsAdminCtx)
	siteRetrieve, err := h.services.Sites.GetSiteByID(c.Request.Context(), isAdmin, objID, siteID)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &siteRetrieve)
}

// @Summary Delete site
// @Security ApiAuth
// @Tags sites
// @Description delete site
// @Accept  json
// @Produce  json
// @Param site_id path string true "site id"
// @Success 204
// @Failure 400,403 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/sites/{site_id}/ [delete]
func (h *APIHandler) siteDelete(c *gin.Context) {
	siteID, err := parseIdFromPath(c, "site_id")
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userID := c.GetString(userIDCtx)
	objID, _ := utils.ConverIDtoObjectId(userID)
	isAdmin := c.GetBool(userIsAdminCtx)
	err = h.services.Sites.DeleteSite(c.Request.Context(), isAdmin, objID, siteID)
	if err != nil {
		newResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// @Summary List of sites
// @Security ApiAuth
// @Tags sites
// @Description list of sites
// @Accept  json
// @Produce  json
// @Param offset query int false "offset" default(0)
// @Param limit query int false "limit" default(100)
// @Param search query string false "search"
// @Success 200 {object} sitesResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/sites/ [get]
func (s *APIHandler) siteList(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	search := queryParams.Get("search")
	limit, offset := getLimitOffsetFromQueryParams(c)
	userID := c.GetString(userIDCtx)
	objID, _ := utils.ConverIDtoObjectId(userID)
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
	c.JSON(http.StatusOK, sitesResponse{Data: data, Count: count})
}

// @Summary Update Site
// @Security ApiAuth
// @Tags sites
// @Description update site
// @Accept  json
// @Produce  json
// @Param site_id path string true "site id"
// @Param input body dto.UpdateSiteDTO true "site info"
// @Success 200 {object} dto.SiteRetrieveDTO
// @Failure 400,403 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/sites/{site_id}/ [put]
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
	objID, _ := utils.ConverIDtoObjectId(userID)
	isAdmin := c.GetBool(userIsAdminCtx)
	res, err := s.services.Sites.UpdateSite(c.Request.Context(), isAdmin, objID, siteID, &siteDTO)
	if err != nil {
		newResponse(c, http.StatusForbidden, core.ErrSiteAccessDenied.Error())
		return
	}
	c.JSON(http.StatusOK, res)

}
