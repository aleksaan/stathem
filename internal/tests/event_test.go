package tests

import (
	"testing"

	"github.com/aleksaan/stathem/internal/actions"
	"github.com/aleksaan/stathem/internal/database"
	"github.com/aleksaan/stathem/internal/entities/events"
	"github.com/aleksaan/stathem/internal/requests"
	"github.com/aleksaan/stathem/internal/utils"
)

func TestEventCreating(t *testing.T) {

	tx := database.DB.Begin()
	defer tx.Commit()

	rp, _ := utils.ConvertJsonToStruct[requests.ProcessCreateRequest](jsonRight)
	p, _ := actions.ProcessCreate(ctx, tx, rp)

	re := &requests.EventCreateRequest{
		ProcessToken: p.ProcessToken,
		StepName:     "S1",
		StateName:    "RUN",
	}

	f := &events.EventsFilter{
		IsActualState:        new(true),
		IsSetWhileProcessRun: new(true),
		IsStepNameRight:      new(true),
		IsStateNameRight:     new(true),
		IsStepSequenceRight:  new(true),
		IsStateSequenceRight: new(true),
	}

	//-----------S1 RUN - OK
	eS1Run, err := events.NewEvent(ctx, tx, re, false)
	if err != nil {
		t.Errorf("Create event STEP1 RUN: await nil, got error")
	}

	if !eS1Run.IsMatchFilter(f) {
		t.Errorf("S1 RUN: doesnt match the filter")
	}

	//-----------S2 RUN - isStepSequenceRight = false
	re = &requests.EventCreateRequest{
		ProcessToken: p.ProcessToken,
		StepName:     "S2",
		StateName:    "RUN",
	}

	f = &events.EventsFilter{
		IsActualState:        new(true),
		IsSetWhileProcessRun: new(true),
		IsStepNameRight:      new(true),
		IsStateNameRight:     new(true),
		IsStepSequenceRight:  new(false),
		IsStateSequenceRight: new(true),
	}

	eS2Run, _ := events.NewEvent(ctx, tx, re, false)

	if !eS2Run.IsMatchFilter(f) {
		t.Errorf("S2 RUN: doesnt match the filter")
	}

	//-----------S3 DONE - isStepSequenceRight = false + isStateSequenceRight = false
	re = &requests.EventCreateRequest{
		ProcessToken: p.ProcessToken,
		StepName:     "S3",
		StateName:    "DONE",
	}

	f = &events.EventsFilter{
		IsActualState:        new(true),
		IsSetWhileProcessRun: new(true),
		IsStepNameRight:      new(true),
		IsStateNameRight:     new(true),
		IsStepSequenceRight:  new(false),
		IsStateSequenceRight: new(false),
	}

	eS3Done, _ := events.NewEvent(ctx, tx, re, false)

	if !eS3Done.IsMatchFilter(f) {
		t.Errorf("S3 DONE: doesnt match the filter")
	}

	//-----------S1 DONE - OK
	re = &requests.EventCreateRequest{
		ProcessToken: p.ProcessToken,
		StepName:     "S1",
		StateName:    "DONE",
	}

	f = &events.EventsFilter{
		IsActualState:        new(true),
		IsSetWhileProcessRun: new(true),
		IsStepNameRight:      new(true),
		IsStateNameRight:     new(true),
		IsStepSequenceRight:  new(true),
		IsStateSequenceRight: new(true),
	}

	eS1Done, _ := events.NewEvent(ctx, tx, re, false)

	if !eS1Done.IsMatchFilter(f) {
		t.Errorf("S1 DONE: doesnt match the filter")
	}

	//-----------S3 RUN - IsStateSequenceRight=false
	re = &requests.EventCreateRequest{
		ProcessToken: p.ProcessToken,
		StepName:     "S3",
		StateName:    "RUN",
	}

	f = &events.EventsFilter{
		IsActualState:        new(true),
		IsSetWhileProcessRun: new(true),
		IsStepNameRight:      new(true),
		IsStateNameRight:     new(true),
		IsStepSequenceRight:  new(true),
		IsStateSequenceRight: new(false),
	}

	eS3Run, _ := events.NewEvent(ctx, tx, re, false)

	if !eS3Run.IsMatchFilter(f) {
		t.Errorf("S3 RUN: doesnt match the filter")
	}

	//-----------S6 RUN - IsStepNameRight=false
	re = &requests.EventCreateRequest{
		ProcessToken: p.ProcessToken,
		StepName:     "S6",
		StateName:    "RUN",
	}

	f = &events.EventsFilter{
		IsActualState:        new(true),
		IsSetWhileProcessRun: new(true),
		IsStepNameRight:      new(false),
		IsStateNameRight:     new(true),
		IsStepSequenceRight:  new(true),
		IsStateSequenceRight: new(true),
	}

	eS6Run, _ := events.NewEvent(ctx, tx, re, false)

	if !eS6Run.IsMatchFilter(f) {
		t.Errorf("S6 RUN: doesnt match the filter")
	}

	//-----------S3 LOAD - IsStateNameRight=false  IsStateSequenceRight=false
	re = &requests.EventCreateRequest{
		ProcessToken: p.ProcessToken,
		StepName:     "S3",
		StateName:    "LOAD",
	}

	f = &events.EventsFilter{
		IsActualState:        new(true),
		IsSetWhileProcessRun: new(true),
		IsStepNameRight:      new(true),
		IsStateNameRight:     new(false),
		IsStepSequenceRight:  new(true),
		IsStateSequenceRight: new(false),
	}

	eS3Load, _ := events.NewEvent(ctx, tx, re, false)

	if !eS3Load.IsMatchFilter(f) {
		t.Errorf("S3 LOAD: doesnt match the filter")
	}

	//-----------S3 ERROR - IsStateNameRight=false  IsStateSequenceRight=false
	re = &requests.EventCreateRequest{
		ProcessToken: p.ProcessToken,
		StepName:     "S3",
		StateName:    "ERROR",
	}

	f = &events.EventsFilter{
		IsActualState:        new(true),
		IsSetWhileProcessRun: new(true),
		IsStepNameRight:      new(true),
		IsStateNameRight:     new(true),
		IsStepSequenceRight:  new(true),
		IsStateSequenceRight: new(false),
	}

	eS3Error, _ := events.NewEvent(ctx, tx, re, false)

	if !eS3Error.IsMatchFilter(f) {
		t.Errorf("S3 ERROR: doesnt match the filter")
	}

}
