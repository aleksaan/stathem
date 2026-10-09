package tests

import (
	"testing"

	"github.com/aleksaan/stathem/internal/config"
	"github.com/aleksaan/stathem/internal/database"
	"github.com/aleksaan/stathem/internal/entities/processes"
	"github.com/aleksaan/stathem/internal/requests"
	"github.com/aleksaan/stathem/internal/utils"
)

var jsonRight = []byte(`
    {
        "model_name": "5 STEPS MODEL",
		"process_params" : {
			"check_steps_names": true,
			"check_states_names": true,
			"check_steps_sequence":true,
			"check_states_sequence":true,
			"check_steps_repeating":true,
			"check_states_repeating":true
			}
    }
`)

var jsonEmptyModel = []byte(`
    {
        "model_name": ""
    }
`)

var jsonWrong = []byte(`
    {
        "X"
    }
`)

func TestProcessValidating(t *testing.T) {
	config.LoadConfig(pathConfig)

	tx := database.DB.Begin()
	defer tx.Rollback()

	// right example
	pr1, _ := utils.ConvertJsonToStruct[requests.ProcessCreateRequest](jsonRight)

	if _, err := processes.New(ctx, tx, pr1); err != nil {
		t.Errorf("Create model from 'jsonRight': await nil, got error")
	}

	// check empty model name
	pr2, _ := utils.ConvertJsonToStruct[requests.ProcessCreateRequest](jsonEmptyModel)

	if _, err := processes.New(ctx, tx, pr2); err == nil {
		t.Errorf("Check process 'jsonEmptyModel' with empty model name: await error, got nil")
	}

	// check wrong json
	if _, err := utils.ConvertJsonToStruct[requests.ProcessCreateRequest](jsonWrong); err == nil {
		t.Errorf("Check process wasnt created'jsonWrong': await error, got nil")
	}
}

func TestProcessFinish(t *testing.T) {

	tx := database.DB.Begin()
	defer tx.Commit()

	pr1, _ := utils.ConvertJsonToStruct[requests.ProcessCreateRequest](jsonRight)

	pdb, err := processes.New(ctx, tx, pr1)

	//errors checking
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if !(pdb.ProcessID > 0) {
		t.Errorf("process wasn't saved in db: ProcessID is null")
	}

	//finishing process (check if it finished later)
	pdb.Finish(ctx, tx)

	//getting created process from db and compare their tokens
	pdb2, err := processes.GetByToken(tx, pdb.ProcessToken, false)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	//check if process token is as we expected
	if pdb.ProcessToken != pdb2.ProcessToken {
		t.Errorf("getting process from db has failed")
	}

	//check if finish state of process is as we expected
	if !pdb2.IsFinished() {
		t.Errorf("expected that process must be finished, but it isnt")
	}

	//trying to get wrong token
	_, err = processes.GetByToken(tx, "X", false)
	if err == nil {
		t.Errorf("Trying to get process with wrong token: await error, got nil")
	}

}
