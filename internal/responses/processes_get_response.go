package responses

import (
	"encoding/json"
	"time"
)

type ProcessesGetResponse struct {
	ProcessesGetResponseInfo []*ProcessesGetResponseInfo
}

type ProcessesGetResponseInfo struct {
	ModelName         string     `json:"model_name"`
	ProcessToken      string     `json:"process_token"`
	ProcessCreatedAt  *time.Time `json:"process_created_at"`
	ProcessFinishedAt *time.Time `json:"process_finished_at"`
}

// Реализуем кастомный метод сериализации для Parent
func (p ProcessesGetResponse) MarshalJSON() ([]byte, error) {
	if p.ProcessesGetResponseInfo == nil {
		return json.Marshal([]ProcessesGetResponseInfo{})
	}
	return json.Marshal(p.ProcessesGetResponseInfo)
}
