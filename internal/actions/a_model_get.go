package actions

import (
	"context"

	"github.com/aleksaan/stathem/internal/entities/models"
	"github.com/aleksaan/stathem/internal/requests"
	"github.com/aleksaan/stathem/internal/responses"
	"github.com/aleksaan/stathem/internal/utils"
	"gorm.io/gorm"
)

func ModelGet(ctx context.Context, tx *gorm.DB, mr *requests.ModelGetRequest) ([]byte, error) {

	var json []byte

	//get model (its active version) from database
	m, err := models.GetActiveByName(tx, mr.ModelName)
	if err != nil {
		return nil, err
	}

	json = utils.ConvertStructToJson(ModelGetResponse(m))

	return json, nil

}

func ModelGetResponse(m *models.Model) *responses.ModelGetResponse {
	mresp := &responses.ModelGetResponse{
		ModelID:            m.ModelID,
		ModelName:          m.ModelName,
		ModelNodes:         m.ModelNodes,
		ModelsEdges:        m.ModelEdges,
		ModelPayload:       m.ModelPayload,
		ModelCreator:       m.ModelCreator,
		ModelCreatedAt:     m.ModelCreatedAt,
		ModelDeactivator:   m.ModelDeactivator,
		ModelDeactivatedAt: m.ModelDeactivatedAt,
	}
	return mresp
}
