package actions

import (
	"context"

	"github.com/aleksaan/stathem/internal/entities/models"
	"github.com/aleksaan/stathem/internal/responses"
	"github.com/aleksaan/stathem/internal/utils"
	"gorm.io/gorm"
)

func ModelsGet(ctx context.Context, tx *gorm.DB) ([]byte, error) {

	var json []byte

	//get model (its active version) from database
	m, err := models.GetAllActiveModels(tx)
	if err != nil {
		return nil, err
	}

	json = utils.ConvertStructToJson(ModelsGetResponse(m))

	return json, nil

}

func ModelsGetResponse(mm []*models.Model) *responses.ModelsGetResponse {
	resp := &responses.ModelsGetResponse{}

	for _, m := range mm {
		m_resp := ModelGetResponse(m)
		resp.ModelGetResponse = append(resp.ModelGetResponse, m_resp)
	}
	return resp
}
