package requests

import "gorm.io/datatypes"

type ProcessCreateRequest struct {
	ModelName     string         `json:"model_name"`
	ProcessParams datatypes.JSON `json:"process_params"`
	// ProcessRules   datatypes.JSON `json:"process_rules"`
	ProcessPayload datatypes.JSON `json:"process_payload"`
}

// type CreateProcessParamsRequest struct {
// 	CheckStepsNames      bool `json:"check_steps_names"`
// 	CheckStatesNames     bool `json:"check_states_names"`
// 	CheckStepsSequence   bool `json:"check_steps_sequence"`
// 	CheckStatesSequence  bool `json:"check_states_sequence"`
// 	CheckStepsRepeating  bool `json:"check_steps_repeating"`
// 	CheckStatesRepeating bool `json:"check_states_repeating"`
// }
