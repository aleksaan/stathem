package requests

type EventsIsReadyForSetRequest struct {
	ProcessToken string `json:"process_token"`
	StepName     string `json:"step_name"`
}
