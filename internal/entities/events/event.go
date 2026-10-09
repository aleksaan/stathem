package events

import (
	"context"
	"time"

	"github.com/aleksaan/stathem/internal/config"
	"github.com/aleksaan/stathem/internal/entities/processes"
	"github.com/aleksaan/stathem/internal/requests"
	"github.com/aleksaan/stathem/internal/utils"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Event struct {
	EventID            uint `gorm:"primaryKey"`
	ProcessID          uint
	Process            processes.Process
	StepName           string         `gorm:"default:null"`
	StateName          string         `gorm:"default:null"`
	EventChecks        EventChecks    `gorm:"serializer:json;type:jsonb"`
	EventRuntimeChecks RuntimeChecks  `gorm:"-"`
	EventPayload       datatypes.JSON `gorm:"type:json"`
	EventCreatedBy     string         `gorm:"not null;default:null"`
	EventCreatedAt     *time.Time     `gorm:"autoCreateTime;not null;default:null"`
}

type RuntimeChecks struct {
	IsActual *bool
}

func (Event) TableName() string {
	return config.Config.DbSchema + ".events"
}

func NewEvent(ctx context.Context, tx *gorm.DB, er *requests.EventCreateRequest, isInit bool) (event *Event, err error) {
	currentTime := time.Now()

	p, err := processes.GetByToken(tx, er.ProcessToken, true)

	//create event
	event = &Event{
		ProcessID:      p.ProcessID,
		StepName:       er.StepName,
		StateName:      er.StateName,
		EventPayload:   er.Payload,
		EventCreatedBy: utils.GetUserNameFromContext(ctx),
		EventCreatedAt: &currentTime,
	}

	event.DoChecks(tx, p, isInit)

	event.EventRuntimeChecks.IsActual = new(true)

	//save event into database
	if err := tx.Create(&event); err.Error != nil {
		return nil, err.Error
	}

	if err := tx.Preload("Process").First(&event, event.EventID).Error; err != nil {
		return nil, err
	}

	return event, nil
}
