package responses

import (
	"time"

	"gorm.io/datatypes"
)

type EventCreateResponse struct {
	ProcessToken       string         `json:"process_token"`
	StepName           string         `json:"step_name"`
	StateName          string         `json:"state_name"`
	EventChecksResults datatypes.JSON `json:"checks_results"`
	EventCreatedBy     string         `json:"created_by"`
	EventCreatedAt     *time.Time     `json:"created_at"`
}
