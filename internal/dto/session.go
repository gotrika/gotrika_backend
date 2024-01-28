package dto

type SessionDTO struct {
	SiteID              string  `json:"site_id"`
	SessionID           string  `json:"session_id"`
	ServerTimestamp     int     `json:"server_timestamp"`
	ClientTimestamp     int     `json:"client_timestamp"`
	ClientTimezone      string  `json:"client_timezone"`
	Duration            int     `json:"duration"`
	VisitorID           string  `json:"visitor_id"`
	UTMSource           string  `json:"utm_source"`
	UTMMedium           string  `json:"utm_medium"`
	UTMCampaign         string  `json:"utm_campaign"`
	UTMContent          string  `json:"utm_content"`
	UTMTerm             string  `json:"utm_term"`
	UTMReferrer         string  `json:"utm_referrer"`
	Referrer            string  `json:"referrer"`
	IPAdress            string  `json:"ip_address"`
	CountryShort        string  `json:"country_short"`
	CountryLong         string  `json:"country_long"`
	RegionName          string  `json:"region_name"`
	City                string  `json:"city"`
	Longtitude          float64 `json:"longtitude"`
	Latitude            float64 `json:"latitude"`
	OSFamily            string  `json:"os_family"`
	OSName              string  `json:"os_name"`
	Browser             string  `json:"browser"`
	BrowserFullVersion  string  `json:"browser_full_version"`
	BrowserMinorVersion string  `json:"browser_minor_version"`
	BrowserMajorVersion string  `json:"browser_major_version"`
	ClientWindowWidth   int     `json:"client_window_width"`
	ClientWindowHeight  int     `json:"client_window_height"`
	Device              string  `json:"device"`
}
