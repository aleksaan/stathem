package responses

import (
	"time"

	"gorm.io/datatypes"
)

type ProcessGetResponse struct {
	ModelName         string         `json:"model_name"`
	ProcessToken      string         `json:"process_token"`
	ProcessParams     datatypes.JSON `json:"process_params"`
	ProcessPayload    datatypes.JSON `json:"process_payload"`
	ProcessCreatedBy  string         `json:"process_created_by"`
	ProcessCreatedAt  *time.Time     `json:"process_created_at"`
	ProcessFinishedBy string         `json:"process_finished_by"`
	ProcessFinishedAt *time.Time     `json:"process_finished_at"`
}
