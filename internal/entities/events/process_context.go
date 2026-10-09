package events

import (
	"github.com/aleksaan/stathem/internal/entities/models"
	"github.com/aleksaan/stathem/internal/entities/processes"
	"gorm.io/gorm"
)

type ProcessContext struct {
	Model   *models.Model
	Process *processes.Process
	Events  *Events
}

// get DBContext which contains information about process, its model and its events
func NewProcessContext(tx *gorm.DB, processToken string) (*ProcessContext, error) {
	prCtx := &ProcessContext{}

	//get process
	p, err := processes.GetByToken(tx, processToken, false)

	if err != nil {
		return nil, err
	}

	prCtx.Process = p

	//get model
	m, err := models.GetById(tx, p.ModelID)

	if err != nil {
		return nil, err
	}

	prCtx.Model = m

	//get events
	e, err := GetEvents(tx, p)

	if err != nil {
		return nil, err
	}

	prCtx.Events = e

	return prCtx, nil
}
