package responses

import (
	"encoding/json"
	"time"

	"gorm.io/datatypes"
)

type EventsGetResponse struct {
	Events []*EventGetResponse
}

type EventGetResponse struct {
	ProcessToken       string         `json:"process_token"`
	StepName           string         `json:"step_name"`
	StateName          string         `json:"state_name"`
	EventChecksResults datatypes.JSON `json:"checks_results"`
	EventCreatedBy     string         `json:"created_by"`
	EventCreatedAt     *time.Time     `json:"created_at"`
}

func (e EventsGetResponse) MarshalJSON() ([]byte, error) {
	if e.Events == nil {
		return json.Marshal([]*EventGetResponse{}) // Возвратит [] вместо null
	}
	return json.Marshal(e.Events)
}

// UnmarshalJSON позволяет правильно распарсить чистый JSON-массив обратно в структуру
func (e *EventsGetResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &e.Events)
}
