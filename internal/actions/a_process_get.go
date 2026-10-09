package actions

import (
	"context"

	"github.com/aleksaan/stathem/internal/entities/processes"
	"github.com/aleksaan/stathem/internal/requests"
	"github.com/aleksaan/stathem/internal/responses"
	"github.com/aleksaan/stathem/internal/utils"
	"gorm.io/gorm"
)

func ProcessGet(ctx context.Context, tx *gorm.DB, process *requests.ProcessGetRequest) ([]byte, error) {

	//get process from database
	p, err := processes.GetByToken(tx, process.ProcessToken, false)
	if err != nil {
		return nil, err
	}

	json := utils.ConvertStructToJson(ProcessGetResponse(p))
	return json, nil

}

func ProcessGetResponse(process *processes.Process) *responses.ProcessGetResponse {
	pr := &responses.ProcessGetResponse{
		ModelName:         process.Model.ModelName,
		ProcessToken:      process.ProcessToken,
		ProcessParams:     process.ProcessParams,
		ProcessPayload:    process.ProcessPayload,
		ProcessCreatedBy:  process.ProcessCreatedBy,
		ProcessCreatedAt:  process.ProcessCreatedAt,
		ProcessFinishedBy: process.ProcessFinishedBy,
		ProcessFinishedAt: process.ProcessFinishedAt,
	}
	return pr
}
