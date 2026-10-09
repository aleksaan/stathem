package checks

import (
	"github.com/aleksaan/stathem/internal/entities/states"
	"github.com/aleksaan/stathem/internal/requests"
	"gorm.io/gorm"
)

type Checks struct {
	IsStepNameRight      *bool `json:"is_step_name_right"`
	IsStateNameRight     *bool `json:"is_state_name_right"`
	IsStepSequenceRight  *bool `json:"is_step_sequence_right"`
	IsStateSequenceRight *bool `json:"is_state_sequence_right"`
	IsSetWhileProcessRun *bool `json:"is_set_while_process_run"`
}

func New(tx *gorm.DB, e *requests.EventCreateRequest) (*Checks, error) {
	pctx, err := NewProcessContext(tx, e.ProcessToken, true)
	if err != nil {
		return nil, err
	}

	c := &Checks{}
	c.runChecks(pctx, e)
	return c, nil
}

func (c *Checks) runChecks(pctx *ProcessContext, e *requests.EventCreateRequest) {

	c.IsStepNameRight = checkStepNameIsValid(pctx, e.StepName)
	c.IsStateNameRight = checkStateNameIsValid(e.StateName)
	c.IsStepSequenceRight = checkStepSequence(pctx, e)
	c.IsStateSequenceRight = checkStateSequence(pctx, e)
	c.IsSetWhileProcessRun = checkIsSetWhileProcessRun(pctx)

}

//--------------------------

func checkStepNameIsValid(pctx *ProcessContext, stepName string) bool {
	for _, v := range pctx.Model.ModelNodes {
		if v == stepName {
			return true
		}
	}
	return false
}

//--------------------------

func checkStateNameIsValid(stateName string) bool {
	return states.CheckStateIsValid(stateName)
}

//--------------------------

// func сheckStepsRepeating(pctx *ProcessContext, e *requests.EventCreateRequest) bool {
// 	for _, ev := range pctx.Events.Events {
// 		if e.StepName == ev.StepName && e.StateName == ev.StateName {
// 			return false
// 		}
// 	}
// 	return true
// }

//--------------------------

func checkStateSequence(pctx *ProcessContext, e *requests.EventCreateRequest) bool {
	lastEvent := pctx.Events.GetLatetsEvent()
	res := states.CheckStateSequence(e.StateName, lastEvent.StateName)

	return res
}

//----------------------------

func checkStepSequence(pctx *ProcessContext, e *requests.EventCreateRequest) bool {
	previousNodes := pctx.Model.GetPreviousNodes(e.StepName)
	count := 0
	for _, node := range previousNodes {
		e := pctx.Events.GetEventByStepName(node)
		if e != nil {
			if states.CheckStateIsFinishState(e.StateName) {
				count++
			}
		}
	}
	if count != len(previousNodes) {
		return false
	}
	return true
}
