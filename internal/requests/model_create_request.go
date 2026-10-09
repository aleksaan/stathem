package requests

import "gorm.io/datatypes"

type ModelCreateRequest struct {
	ModelName    string              `json:"model_name"`
	ModelNodes   []string            `json:"model_nodes,omitempty"`
	ModelEdges   []map[string]string `json:"model_edges,omitempty"`
	ModelPayload datatypes.JSON      `json:"model_payload,omitempty"`
}
