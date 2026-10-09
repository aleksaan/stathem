package actions

import (
	"context"

	"github.com/aleksaan/stathem/internal/entities/models"
	"github.com/aleksaan/stathem/internal/requests"
	"gorm.io/gorm"
)

// ModelCreate - create model from api data and save it to db
func ModelCreate(ctx context.Context, tx *gorm.DB, m *requests.ModelCreateRequest) error {

	//creating model and save to database
	_, err := models.New(ctx, tx, m)
	if err != nil {
		return err
	}

	return nil

}
