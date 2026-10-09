package processes

import (
	"github.com/aleksaan/stathem/internal/entities/models"

	"gorm.io/gorm"
)

func GetByModelName(tx *gorm.DB, modelName string) ([]*Process, error) {

	var processes []*Process

	m, err := models.GetActiveByName(tx, modelName)

	if err != nil {
		return nil, err
	}

	res := tx.Preload("Model").Where(&Process{ModelID: m.ModelID}).Find(&processes)

	if res.Error != nil {
		return nil, res.Error
	}

	return processes, nil
}
