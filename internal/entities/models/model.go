package models

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/aleksaan/stathem/internal/config"
	apperrors "github.com/aleksaan/stathem/internal/errors"
	"github.com/aleksaan/stathem/internal/requests"
	"github.com/aleksaan/stathem/internal/utils"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Model struct {
	ModelID            uint                `gorm:"primaryKey"`
	ModelName          string              `gorm:"not null"`
	ModelNodes         []string            `gorm:"serializer:json"`
	ModelEdges         []map[string]string `gorm:"serializer:json"`
	ModelPayload       datatypes.JSON      `gorm:"type:json"`
	ModelCreator       string              `gorm:"not null"`
	ModelCreatedAt     *time.Time          `gorm:"autoCreateTime; not null"`
	ModelDeactivator   string              `gorm:"not null"`
	ModelDeactivatedAt *time.Time
}

func (Model) TableName() string {
	return config.Config.DbSchema + ".models"
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	err = DeactivateSameModels(tx, m)

	if err != nil {
		return err
	}
	return
}

// creating new model
func New(ctx context.Context, tx *gorm.DB, mr *requests.ModelCreateRequest) (*Model, error) {

	model := &Model{
		ModelName:    mr.ModelName,
		ModelNodes:   mr.ModelNodes,
		ModelEdges:   mr.ModelEdges,
		ModelPayload: mr.ModelPayload,
		ModelCreator: utils.GetUserNameFromContext(ctx)}

	if err := model.Validate(); err != nil {
		return nil, err
	}

	if err := tx.Create(&model); err.Error != nil {
		return nil, err.Error
	}

	return model, nil
}

// get model by ModelID
func GetById(tx *gorm.DB, modelID uint) (*Model, error) {
	model := &Model{}

	result := tx.
		Where(&Model{ModelID: modelID}).
		First(&model)

	if result.Error != nil {
		return nil, apperrors.ErrModelNotFound
	}

	return model, nil
}

// get active model by ModelName
func GetActiveByName(tx *gorm.DB, modelName string) (*Model, error) {
	model := &Model{}

	err := tx.
		Where("model_name = ?", modelName).
		Where("model_deactivated_at is null").
		First(&model).Error

	if err != nil {
		return nil, apperrors.ErrModelNotFound
	}

	return model, nil
}

func GetAllActiveModels(tx *gorm.DB) ([]*Model, error) {
	var models []*Model

	err := tx.
		Where("model_deactivated_at is null").
		Find(&models).Error

	if err != nil {
		return nil, apperrors.ErrModelNotFound
	}

	return models, nil
}

func DeactivateSameModels(tx *gorm.DB, m *Model) error {
	currentTime := time.Now()

	result := tx.Model(&Model{}).
		Where(&Model{ModelName: m.ModelName}).
		Where("model_deactivated_at is null").
		Updates(Model{ModelDeactivatedAt: &currentTime})

	if result.Error != nil {
		return result.Error
	}

	slog.Debug("Deactivating previous active version of this model")
	slog.Debug(fmt.Sprintf("Number of deactivated versions: %d", result.RowsAffected))

	return nil
}
