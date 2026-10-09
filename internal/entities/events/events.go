package events

import (
	"slices"

	"github.com/aleksaan/stathem/internal/entities/processes"
	"gorm.io/gorm"
)

type Events struct {
	Events       []*Event
	LatestEvents map[string]*Event
}

// func (ee *Events) GetEventsByStepName(stepName string, isActual *bool) *Events {

// 	res := &Events{}

// 	for _, e := range ee.Events {
// 		if stepName == e.StepName {
// 			if isActual == nil || e.EventRuntimeChecks.IsActual == isActual {
// 				res.Events = append(res.Events, e)
// 			}
// 		}
// 	}

// 	return res
// }

// func (ee *Events) GetActualState(stepName string) (*Event, error) {
// 	nn := ee.GetEventsByStepName(stepName, new(true))
// 	for _, e := range nn.Events {
// 		if e.EventRuntimeChecks.IsActual == new(true) {
// 			return e, nil
// 		}
// 	}
// 	return nil, apperrors.ErrActualStateNotFound
// }

func GetEvents(tx *gorm.DB, process *processes.Process) (*Events, error) {

	events := &Events{}

	err := tx.Preload("Process").Where("process_id=?", process.ProcessID).Find(&events.Events).Error

	if err != nil {
		return nil, err
	}

	events.findLatests()
	events.setIsActual()

	return events, nil
}

// findLatests - preparing map of latests events for futher using
func (ee *Events) findLatests() {

	ee.LatestEvents = make(map[string]*Event)

	for _, comparedValue := range ee.Events {
		latestValue, exists := ee.LatestEvents[comparedValue.StepName]
		if !exists {
			ee.LatestEvents[comparedValue.StepName] = comparedValue
		} else {
			if latestValue.EventCreatedAt.Before(*comparedValue.EventCreatedAt) {
				ee.LatestEvents[comparedValue.StepName] = comparedValue
			}
		}
	}
}

// setIsActual - set if state is actual (latest) on this moment
func (ee *Events) setIsActual() {
	for _, e := range ee.Events {
		l, exists := ee.LatestEvents[e.StepName]
		e.EventRuntimeChecks.IsActual = new(false)
		if exists {
			if l.EventID == e.EventID {
				e.EventRuntimeChecks.IsActual = new(true)
			}
		}
	}
}

func (ee *Events) GetPrevious(tx *gorm.DB, process *processes.Process, stepName string) (*Events, error) {

	prevSteps := process.Model.GetPreviousNodes(stepName)

	prevEvents := &Events{}

	for _, e := range ee.Events {
		if slices.Contains(prevSteps, e.StepName) {
			prevEvents.Events = append(prevEvents.Events, e)
		}
	}

	return prevEvents, nil
}
