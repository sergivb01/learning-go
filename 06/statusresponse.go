package main

// StatusResponse is defined by the service status
type StatusResponse struct {
	Page                  Page          `json:"page"`
	Components            []Component   `json:"components"`
	Incidents             []interface{} `json:"incidents"`
	ScheduledMaintenances []interface{} `json:"scheduled_maintenances"`
	Status                Status        `json:"status"`
}

// Component defines the basic strucuture of a service check
type Component struct {
	Status             string      `json:"status"`
	Name               string      `json:"name"`
	CreatedAt          string      `json:"created_at"`
	UpdatedAt          string      `json:"updated_at"`
	Position           int64       `json:"position"`
	Description        *string     `json:"description"`
	Showcase           bool        `json:"showcase"`
	ID                 string      `json:"id"`
	GroupID            interface{} `json:"group_id"`
	PageID             string      `json:"page_id"`
	Group              bool        `json:"group"`
	OnlyShowIfDegraded bool        `json:"only_show_if_degraded"`
}

type Page struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	TimeZone  string `json:"time_zone"`
	UpdatedAt string `json:"updated_at"`
}

type Status struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}
