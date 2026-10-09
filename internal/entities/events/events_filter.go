package events

import (
	"github.com/aleksaan/stathem/internal/entities/processes"
	"gorm.io/gorm"
)

type EventsFilter struct {
	IsActualState        *bool
	IsReadyToSet         *bool
	IsStepNameRight      *bool
	IsStateNameRight     *bool
	IsStepSequenceRight  *bool
	IsStateSequenceRight *bool
	IsSetWhileProcessRun *bool
	IsChildStepName      *string
}

func (e *Event) IsMatchFilter(f *EventsFilter) bool {
	res := true

	if f.IsActualState != nil {
		if *e.EventRuntimeChecks.IsActual != *f.IsActualState {
			res = false
		}
	}
	if f.IsSetWhileProcessRun != nil {
		if *e.EventChecks.IsSetWhileProcessRun != *f.IsSetWhileProcessRun {
			res = false
		}
	}
	if f.IsStepNameRight != nil {
		if *e.EventChecks.IsStepNameRight != *f.IsStepNameRight {
			res = false
		}
	}
	if f.IsStepSequenceRight != nil {
		if *e.EventChecks.IsStepSequenceRight != *f.IsStepSequenceRight {
			res = false
		}
	}
	if f.IsStateNameRight != nil {
		if *e.EventChecks.IsStateNameRight != *f.IsStateNameRight {
			res = false
		}
	}
	if f.IsStateSequenceRight != nil {
		if *e.EventChecks.IsStateSequenceRight != *f.IsStateSequenceRight {
			res = false
		}
	}

	return res
}

func (ee *Events) GetByFilter(tx *gorm.DB, p *processes.Process, f *EventsFilter) *Events {
	filtered := &Events{}

	parents := &Events{}

	//get only parents if IsChildStepName is set
	if f.IsChildStepName != nil {
		parents, _ = ee.GetPrevious(tx, p, *f.IsChildStepName)
	} else {
		parents = ee
	}

	//filter by other filters
	for _, e := range parents.Events {
		if e.IsMatchFilter(f) {
			filtered.Events = append(filtered.Events, e)
		}
	}

	return filtered
}
