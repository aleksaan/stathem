package actions

import (
	"context"

	"github.com/aleksaan/stathem/internal/entities/events"
	"github.com/aleksaan/stathem/internal/requests"
	"github.com/aleksaan/stathem/internal/responses"
	"github.com/aleksaan/stathem/internal/utils"
	"gorm.io/gorm"
)

func EventCreate(ctx context.Context, tx *gorm.DB, er *requests.EventCreateRequest) ([]byte, error) {

	//create event
	e, err := events.NewEvent(ctx, tx, er, false)
	if err != nil {
		return nil, err
	}

	json := utils.ConvertStructToJson(EventCreateResponse(e))
	return json, nil

}

func EventCreateResponse(event *events.Event) *responses.EventCreateResponse {

	e := &responses.EventCreateResponse{
		ProcessToken:       event.Process.ProcessToken,
		StepName:           event.StepName,
		StateName:          event.StateName,
		EventChecksResults: utils.ConvertStructToJson(event.EventChecks),
		EventCreatedBy:     event.EventCreatedBy,
		EventCreatedAt:     event.EventCreatedAt,
	}

	return e
}
