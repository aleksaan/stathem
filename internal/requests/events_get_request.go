package requests

type EventsGetRequest struct {
	ProcessToken string                 `json:"process_token"`
	Filter       EventsGetRequestFilter `json:"filter"`
}

type EventsGetRequestFilter struct {
	IsActualOnly         *bool   `json:"is_actual_only,omitempty"`
	IsReadyToSetOnly     *bool   `json:"is_ready_to_set_only,omitempty"`
	IsSetWhileProcessRun *bool   `json:"is_set_while_process_run,omitempty"`
	IsStepNameRight      *bool   `json:"is_step_name_right,omitempty"`
	IsStateNameRight     *bool   `json:"is_state_name_right,omitempty"`
	IsStepSequenceRight  *bool   `json:"is_step_sequence_right,omitempty"`
	IsStateSequenceRight *bool   `json:"is_state_sequence_right,omitempty"`
	IsChildStepName      *string `json:"is_child_step_name,omitempty"`
}
