package collect

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"github.com/gotrika/gotrika_backend/internal/transport/http/handler/utils"
)

func (h *CollectHandler) initTrackerHandlers(collect *gin.RouterGroup) {
	{
		collect.POST("/", h.CollectData)
	}

}

type CollectRequest struct {
	SiteID         string          `json:"site_id"`
	HashID         string          `json:"hash_id"`
	Timestamp      int             `json:"timestamp"`
	Type           string          `json:"type"`
	RawTrackerData json.RawMessage `json:"tracker_data"`
}

func (h *CollectHandler) CollectData(c *gin.Context) {
	var inp CollectRequest
	realIP := c.ClientIP()
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	siteID, err := utils.ConverIDtoObjectId(inp.SiteID)
	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid site id")
		return
	}
	trackerData, err := inp.RawTrackerData.MarshalJSON()
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	addDTO := dto.TrackerDataDTO{
		SiteID:      siteID,
		HashID:      inp.HashID,
		Timestamp:   time.Unix(int64(inp.Timestamp), 0),
		Type:        inp.Type,
		TrackerData: trackerData,
		RealIP:      realIP,
	}
	err = h.services.TrackerService.SaveRawTrackerData(c.Request.Context(), &addDTO)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	resp := trackerResponse{Success: true}
	c.JSON(http.StatusOK, &resp)
}
