package actions

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aleksaan/stathem/internal/config"
	"github.com/aleksaan/stathem/internal/entities/models"
	"github.com/aleksaan/stathem/internal/requests"
	"github.com/aleksaan/stathem/internal/utils"
	"gorm.io/gorm"
)

func CreateDefaultModels(ctx context.Context, tx *gorm.DB) error {
	if config.Config.DbRecreateTables {
		slog.Info("Creating default models")

		for _, jsonb := range models.JsonDefaultModels {
			rm, err := utils.ConvertJsonToStruct[requests.ModelCreateRequest](jsonb)
			if err != nil {
				slog.Error(err.Error())
				return err
			}

			_, err = models.New(ctx, tx, rm)
			if err != nil {
				slog.Error(err.Error())
				return err
			}

			slog.Info(fmt.Sprintf("--- Model %s was created", rm.ModelName))
		}

		slog.Info("Creating default models... Done")

	}

	return nil
}
