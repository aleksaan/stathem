package responses

import (
	"time"

	"gorm.io/datatypes"
)

type ModelGetResponse struct {
	ModelID            uint                `json:"model_id"`
	ModelName          string              `json:"model_name"`
	ModelNodes         []string            `json:"model_nodes"`
	ModelsEdges        []map[string]string `json:"model_edges"`
	ModelPayload       datatypes.JSON      `json:"model_payload"`
	ModelCreator       string              `json:"model_creator"`
	ModelCreatedAt     *time.Time          `json:"model_created_at"`
	ModelDeactivator   string              `json:"model_deactivator"`
	ModelDeactivatedAt *time.Time          `json:"model_deactivated_at"`
}
