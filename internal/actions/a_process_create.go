package actions

import (
	"context"

	"github.com/aleksaan/stathem/internal/entities/events"
	"github.com/aleksaan/stathem/internal/entities/models"
	"github.com/aleksaan/stathem/internal/entities/processes"
	"github.com/aleksaan/stathem/internal/entities/states"
	"github.com/aleksaan/stathem/internal/requests"
	"github.com/aleksaan/stathem/internal/responses"
	"gorm.io/gorm"
)

// CreateProcess - create process from api data, save it to db and set all its steps into NONE state
func ProcessCreate(ctx context.Context, tx *gorm.DB, pr *requests.ProcessCreateRequest) (*processes.Process, error) {

	var p *processes.Process

	//creating new process in database
	p, err := processes.New(ctx, tx, pr)
	if err != nil {
		return nil, err
	}

	//creating new process in database
	err = initEventsForNewProcess(ctx, tx, p)
	if err != nil {
		return nil, err
	}

	return p, nil

	// json := utils.ConvertStructToJson(ProcessCreateResponse(p))
	// return json, nil

}

func initEventsForNewProcess(ctx context.Context, tx *gorm.DB, p *processes.Process) error {
	//get model with nodes
	m, err := models.GetById(tx, p.ModelID)
	if err != nil {
		return err
	}

	//set events into initial state for every node
	for _, n := range m.ModelNodes {
		er := &requests.EventCreateRequest{
			ProcessToken: p.ProcessToken,
			StepName:     n,
			StateName:    states.STATE_NONE,
		}

		_, err := events.NewEvent(ctx, tx, er, true)
		if err != nil {
			return err
		}
	}
	return nil
}

func ProcessCreateResponse(p *processes.Process) *responses.ProcessCreateResponse {
	mresp := &responses.ProcessCreateResponse{
		ProcessToken: p.ProcessToken,
	}
	return mresp
}
