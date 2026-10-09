package database

import (
	"context"
	"log/slog"

	"github.com/aleksaan/stathem/internal/actions"
	"github.com/aleksaan/stathem/internal/config"
	"github.com/aleksaan/stathem/internal/entities/events"
	"github.com/aleksaan/stathem/internal/entities/models"
	"github.com/aleksaan/stathem/internal/entities/processes"
	"github.com/aleksaan/stathem/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB - база данных GORM
var DB *gorm.DB

func InitDBConnection() {
	slog.Info("Database connection")
	ctx := context.Background()
	ctx = utils.InitSystemContext(ctx)

	conn, err := gorm.Open(postgres.Open(config.Config.DbConnectionString), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
		//SkipDefaultTransaction: true,
	})

	if err != nil {
		slog.Error(err.Error())
	}
	DB = conn
	slog.Info("Database connection...Done")

	if config.Config.DbRecreateTables {
		dbTablesMigration()
		actions.CreateDefaultModels(ctx, DB)
	}

}

func dbTablesMigration() {
	tx := DB.Begin()
	defer tx.Commit()

	slog.Info("Droping tables")
	tx.Migrator().DropTable(&models.Model{}, &processes.Process{}, &events.Event{})
	slog.Info("Droping tables... Done")
	slog.Info("Updating tables")
	tx.AutoMigrate(&models.Model{}, &processes.Process{}, &events.Event{})
	slog.Info("Updating tables... Done")
}
