package actions

import (
	"context"

	"github.com/aleksaan/stathem/internal/entities/processes"
	"github.com/aleksaan/stathem/internal/requests"
	"gorm.io/gorm"
)

func ProcessFinish(ctx context.Context, tx *gorm.DB, process *requests.ProcessFinishRequest) error {

	p, err := processes.GetByToken(tx, process.ProcessToken, true)
	if err != nil {
		return err
	}

	err = p.Finish(ctx, tx)

	if err != nil {
		return err
	}

	return nil
}
