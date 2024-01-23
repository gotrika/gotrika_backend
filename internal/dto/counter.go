package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddRawTrackerDataDTO struct {
	SiteID      primitive.ObjectID `json:"site_id"`
	HashID      string             `json:"hash_id"`
	Datetime    time.Time          `json:"datetime"`
	Type        string             `json:"type"`
	TrackerData []byte             `json:"tracked_data"`
}
