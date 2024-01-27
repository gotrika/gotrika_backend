package dto

import "encoding/json"

type ParseTask struct {
	IDS  []string `json:"ids"`
	Type string   `json:"type"`
}

func (t *ParseTask) ToPublish() ([]byte, error) {
	js, err := json.Marshal(t)
	if err != nil {
		return []byte(""), err
	}
	return js, nil
}

type SessionTaskDTO struct {
	SessionID        string `json:"session_id"`
	VisitorID        string `json:"visitor_id"`
	UserAgent        string `json:"user_agent"`
	Location         string `json:"location"`
	ClientTimezone   string `json:"client_timezone"`
	SessionTimestamp int    `json:"session_timestamp"`
	Referrer         string `json:"referrer"`
}

type EventTaskDTO struct {
	SessionID       string `json:"session_id"`
	VisitorID       string `json:"visitor_id"`
	ClassName       string `json:"class_name"`
	Type            string `json:"event_type"`
	ClientTimestamp int    `json:"client_timestamp"`
	Page            string `json:"page"`
	PageTitle       string `json:"page_title"`
	HitURL          string `json:"hit_url"`
}
