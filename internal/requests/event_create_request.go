package requests

import "gorm.io/datatypes"

type EventCreateRequest struct {
	ProcessToken string         `json:"process_token"`
	StepName     string         `json:"step_name"`
	StateName    string         `json:"state_name"`
	Payload      datatypes.JSON `json:"payload"`
}
