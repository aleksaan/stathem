package rules

import (
	"github.com/aleksaan/stathem/internal/entities/states"
	"github.com/aleksaan/stathem/internal/utils"
	"gorm.io/datatypes"
)

type Rules struct {
	ProcessReadyToFinish []string `json:"process_ready_to_finish"`
	StepReadyToSet       []string `json:"steps_ready_to_set"`
}

func New(json datatypes.JSON) (*Rules, error) {

	r, err := utils.ConvertJsonToStruct[Rules](json)

	if err != nil {
		return nil, err
	}

	if len(r.ProcessReadyToFinish) == 0 {
		r.ProcessReadyToFinish = append(r.ProcessReadyToFinish, "*")
	}

	if len(r.StepReadyToSet) == 0 {
		r.ProcessReadyToFinish = append(r.StepReadyToSet, "*")
	}

	for _, v := range r.ProcessReadyToFinish {
		if !(states.CheckStateIsValid(v) || v == "*") {
			return nil, err
		}
	}

	for _, v := range r.StepReadyToSet {
		if !(states.CheckStateIsValid(v) || v == "*") {
			return nil, err
		}
	}

	return r, nil
}
