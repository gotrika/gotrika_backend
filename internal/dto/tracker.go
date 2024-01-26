package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TrackerDataDTO struct {
	SiteID      primitive.ObjectID `json:"site_id"`
	HashID      string             `json:"hash_id"`
	Timestamp   time.Time          `json:"timestamp"`
	Type        string             `json:"type"`
	TrackerData []byte             `json:"tracker_data"`
}
