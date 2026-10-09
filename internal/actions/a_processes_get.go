package actions

import (
	"context"

	"github.com/aleksaan/stathem/internal/entities/processes"
	"github.com/aleksaan/stathem/internal/requests"
	"github.com/aleksaan/stathem/internal/responses"
	"github.com/aleksaan/stathem/internal/utils"
	"gorm.io/gorm"
)

func ProcessesGet(ctx context.Context, tx *gorm.DB, rq *requests.ProcessesGetRequest) ([]byte, error) {

	//get process from database
	p, err := processes.GetByModelName(tx, rq.ModelName)
	if err != nil {
		return nil, err
	}

	json := utils.ConvertStructToJson(ProcessesGetResponse(p))
	return json, nil

}

func ProcessesGetResponse(processes []*processes.Process) *responses.ProcessesGetResponse {

	var res responses.ProcessesGetResponse

	for _, p := range processes {
		pr := &responses.ProcessesGetResponseInfo{
			ModelName:         p.Model.ModelName,
			ProcessToken:      p.ProcessToken,
			ProcessCreatedAt:  p.ProcessCreatedAt,
			ProcessFinishedAt: p.ProcessFinishedAt,
		}
		res.ProcessesGetResponseInfo = append(res.ProcessesGetResponseInfo, pr)
	}
	return &res
}
