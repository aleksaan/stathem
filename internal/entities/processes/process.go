package processes

import (
	"context"
	"errors"
	"time"

	"github.com/aleksaan/stathem/internal/config"
	"github.com/aleksaan/stathem/internal/entities/models"
	apperrors "github.com/aleksaan/stathem/internal/errors"
	"github.com/aleksaan/stathem/internal/requests"
	"github.com/aleksaan/stathem/internal/utils"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Process struct {
	ProcessID     uint `gorm:"primaryKey"`
	ModelID       uint
	Model         models.Model
	ProcessToken  string         `gorm:"unique;not null;default:null"`
	ProcessParams datatypes.JSON `gorm:"type:json"`
	//ProcessRules      datatypes.JSON `gorm:"type:json"`
	ProcessPayload    datatypes.JSON `gorm:"type:json"`
	ProcessCreatedBy  string         `gorm:"not null;default:null"`
	ProcessCreatedAt  *time.Time     `gorm:"autoCreateTime;not null;default:null"`
	ProcessFinishedBy string         `gorm:"default:null"`
	ProcessFinishedAt *time.Time     `gorm:"default:null"`
}

type ProcessRules struct {
	ProcessReadyToFinish []string `json:"process_ready_to_finish"`
	StepReadyToSet       []string `json:"step_ready_to_set"`
}

func (Process) TableName() string {
	return config.Config.DbSchema + ".processes"
}

func New(ctx context.Context, tx *gorm.DB, r *requests.ProcessCreateRequest) (process *Process, err error) {

	currentTime := time.Now()

	m, err := models.GetActiveByName(tx, r.ModelName)
	if err != nil {
		return nil, err
	}

	//create process
	process = &Process{
		ModelID:          m.ModelID,
		ProcessToken:     uuid.New().String(),
		ProcessParams:    r.ProcessParams,
		ProcessPayload:   r.ProcessPayload,
		ProcessCreatedAt: &currentTime,
		ProcessCreatedBy: utils.GetUserNameFromContext(ctx),
	}

	//save process into database
	if err := tx.Preload("Model").Create(&process); err.Error != nil {
		return nil, err.Error
	}

	return process, nil
}

func Validate() {

}

func GetByToken(tx *gorm.DB, token string, forUpdate bool) (process *Process, err error) {
	process = &Process{}

	if token == "" {
		return nil, apperrors.ErrProcessTokenIsEmpty
	}

	query := tx.Preload("Model").Where(&Process{ProcessToken: token}).Clauses(clause.Locking{
		Strength: "UPDATE"})

	if forUpdate {
		query = query.Clauses(clause.Locking{Strength: "UPDATE"})
	}

	res := query.First(&process)

	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, apperrors.ErrProcessNotFound
	}

	if res.Error != nil {
		return nil, apperrors.ErrProcessCannotBeLocked
	}

	return process, nil
}

func (process *Process) Finish(ctx context.Context, tx *gorm.DB) (err error) {
	currentTime := time.Now()

	if !process.IsFinished() {
		res := tx.Where(&process).Updates(Process{ProcessFinishedAt: &currentTime, ProcessFinishedBy: utils.GetUserNameFromContext(ctx)})
		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}

func (process *Process) IsFinished() (res bool) {
	return process.ProcessFinishedAt != nil
}
