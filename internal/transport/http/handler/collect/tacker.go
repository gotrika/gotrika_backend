package collect

import (
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
	SiteID      string `json:"site_id"`
	HashID      string `json:"hash_id"`
	Datetime    int    `json:"datetime"`
	Type        string `json:"type"`
	TrackerData []byte `json:"tracked_data"`
}

func (h *CollectHandler) CollectData(c *gin.Context) {
	var inp CollectRequest
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	siteID, err := utils.ConverIDtoObjectId(inp.SiteID)
	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid site id")
		return
	}
	addDTO := dto.AddRawTrackerDataDTO{
		SiteID:      siteID,
		HashID:      inp.HashID,
		Datetime:    time.Unix(int64(inp.Datetime), 0),
		Type:        inp.Type,
		TrackerData: inp.TrackerData,
	}
	err = h.services.TrackerService.SaveRawTrackerData(c.Request.Context(), &addDTO)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	resp := trackerResponse{Success: true}
	c.JSON(http.StatusOK, &resp)

}
