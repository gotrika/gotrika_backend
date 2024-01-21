package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateSiteDTO struct {
	Name        string `json:"name" binding:"required"`
	URL         string `json:"url" binding:"required"`
	CounterCode string `json:"counter_code,omitempty"`
}

type UpdateSiteDTO struct {
	Name        string   `json:"name"`
	URL         string   `json:"url"`
	AccessUsers []string `json:"access_users"`
}

type SiteRetrieveDTO struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	URL         string   `json:"url"`
	OwnerID     string   `json:"owner_id"`
	CounterCode string   `json:"counter_code"`
	AccessUsers []string `json:"access_users"`
}

type ListSitesDTO struct {
	Limit   int
	Offset  int
	Search  string
	UserID  primitive.ObjectID
	IsAdmin bool
}
