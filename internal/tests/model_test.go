package tests

import (
	"testing"

	"github.com/aleksaan/stathem/internal/database"
	"github.com/aleksaan/stathem/internal/entities/models"
	"github.com/aleksaan/stathem/internal/requests"
	"github.com/aleksaan/stathem/internal/utils"
)

var JsonTestModel = []byte(`
    {
        "model_name": "TEST MODEL",
        "model_nodes": ["S1","S2","S3","S4","S5"],
        "model_edges": [{"S1":"S2"},{"S1":"S3"},{"S1":"S4"},{"S2":"S5"},{"S3":"S5"},{"S4":"S5"}],
		"model_payload": ""
    }
`)

func TestModelCreating(t *testing.T) {

	tx := database.DB.Begin()
	defer tx.Commit()

	//creating test model
	mr, err := utils.ConvertJsonToStruct[requests.ModelCreateRequest](JsonTestModel)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	dbm1, err := models.New(ctx, tx, mr)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if !(dbm1.ModelID > 0) {
		t.Errorf("model wasn't saved in db")
	}

	//creating model with the same name
	dbm2, err := models.New(ctx, tx, mr)
	if err != nil {
		// t.Errorf("expected nil while creating model with the same name, got error")
		t.Errorf("%s", err.Error())
	}

	_, err = models.GetById(tx, dbm1.ModelID)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	gotModel2, err := models.GetActiveByName(tx, dbm2.ModelName)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if gotModel2.ModelName != dbm2.ModelName {
		t.Errorf("wrong logic of finding model by GetActiveByName")
	}
}
