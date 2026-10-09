package actions

import (
	"context"

	"github.com/aleksaan/stathem/internal/entities/events"
	"github.com/aleksaan/stathem/internal/entities/processes"
	"github.com/aleksaan/stathem/internal/requests"
	"github.com/aleksaan/stathem/internal/responses"
	"github.com/aleksaan/stathem/internal/utils"
	"gorm.io/gorm"
)

func StatusesGet(ctx context.Context, tx *gorm.DB, er *requests.EventsGetRequest) ([]byte, error) {

	f := &events.EventsFilter{
		IsActualState:        er.Filter.IsActualOnly,
		IsReadyToSet:         er.Filter.IsReadyToSetOnly,
		IsStepNameRight:      er.Filter.IsStepNameRight,
		IsStepSequenceRight:  er.Filter.IsStepSequenceRight,
		IsStateNameRight:     er.Filter.IsStateNameRight,
		IsStateSequenceRight: er.Filter.IsStateSequenceRight,
		IsChildStepName:      er.Filter.IsChildStepName,
	}

	p, err := processes.GetByToken(tx, er.ProcessToken, false)
	if err != nil {
		return nil, err
	}

	e, err := events.GetEvents(tx, p)
	if err != nil {
		return nil, err
	}

	e = e.GetByFilter(tx, p, f)

	res := EventsGetResponse(e)
	json := utils.ConvertStructToJson(res)
	return json, nil
}

func EventsGetResponse(events *events.Events) *responses.EventsGetResponse {
	resp := &responses.EventsGetResponse{}
	for _, ee := range events.Events {
		e := &responses.EventGetResponse{
			ProcessToken:       ee.Process.ProcessToken,
			StepName:           ee.StepName,
			StateName:          ee.StateName,
			EventChecksResults: utils.ConvertStructToJson(ee.EventChecks),
			EventCreatedBy:     ee.EventCreatedBy,
			EventCreatedAt:     ee.EventCreatedAt,
		}
		resp.Events = append(resp.Events, e)
	}
	return resp
}
