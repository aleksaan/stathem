package events

import (
	"github.com/aleksaan/stathem/internal/entities/processes"
	"github.com/aleksaan/stathem/internal/entities/states"
	"gorm.io/gorm"
)

type EventChecks struct {
	IsStepNameRight      *bool `json:"is_step_name_right"`
	IsStateNameRight     *bool `json:"is_state_name_right"`
	IsStepSequenceRight  *bool `json:"is_step_sequence_right"`
	IsStateSequenceRight *bool `json:"is_state_sequence_right"`
	IsSetWhileProcessRun *bool `json:"is_set_while_process_run"`
}

func (e *Event) DoChecks(tx *gorm.DB, p *processes.Process, isInit bool) error {

	pctx, err := NewProcessContext(tx, p.ProcessToken)
	if err != nil {
		return err
	}

	if isInit {
		e.EventChecks.IsStepNameRight = new(true)
		e.EventChecks.IsStateNameRight = new(true)
		e.EventChecks.IsStepSequenceRight = new(true)
		e.EventChecks.IsStateSequenceRight = new(true)
		e.EventChecks.IsSetWhileProcessRun = new(true)
	} else {
		e.EventChecks.IsStepNameRight = checkStepNameIsValid(pctx, e)
		e.EventChecks.IsStateNameRight = checkStateNameIsValid(e)
		e.EventChecks.IsStepSequenceRight = checkStepSequence(pctx, e)
		e.EventChecks.IsStateSequenceRight = checkStateSequence(pctx, e)
		e.EventChecks.IsSetWhileProcessRun = checkIsSetWhileProcessRun(pctx)
	}

	return nil
}

//--------------------------

func checkIsSetWhileProcessRun(pctx *ProcessContext) *bool {
	return new(!pctx.Process.IsFinished())
}

//--------------------------

func checkStepNameIsValid(pctx *ProcessContext, e *Event) *bool {
	for _, v := range pctx.Model.ModelNodes {
		if v == e.StepName {
			return new(true)
		}
	}
	return new(false)
}

//--------------------------

func checkStateNameIsValid(e *Event) *bool {
	return new(states.CheckStateIsValid(e.StateName))
}

//--------------------------

func checkStateSequence(pctx *ProcessContext, e *Event) *bool {
	actualEvent, exists := pctx.Events.LatestEvents[e.StepName]

	if !exists {
		return new(true)
	}

	return new(states.CheckStateSequence(e.StateName, actualEvent.StateName))
}

//----------------------------

func checkStepSequence(pctx *ProcessContext, e *Event) *bool {
	previousNodes := pctx.Model.GetPreviousNodes(e.StepName)
	count := 0
	for _, node := range previousNodes {
		e, exists := pctx.Events.LatestEvents[node]

		if !exists {
			return new(false)
		}

		if states.CheckStateIsFinishState(e.StateName) {
			count++
		}
	}
	return new(count == len(previousNodes))
}
