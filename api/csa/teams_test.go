package csa_test

import (
	"encoding/json"
	"github.com/fossteams/teams-api/api"
	"github.com/fossteams/teams-api/api/csa"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)
func initTest(t *testing.T) *csa.CSASvc {
	token, err := api.GetTokenFromEnv()
	if err != nil {
		t.Error(err)
	}
	csaSvc, err := csa.NewCSAService(token)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
	return csaSvc
}

func TestGetTeams(t *testing.T){
	csaSvc := initTest(t)
	teams, err := csaSvc.GetTeams()
	
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, teams)
	assert.Greater(t, len(*teams), 0)
}

func TestDecodeTeamsJson(t *testing.T) {
	f, err := os.Open("/home/dvitali/Documents/tmp/teams/api/csa/teams/users/me/response.json")
	if err != nil {
		t.Error(err)
	}
	var meResponse csa.ConversationResponse
	decoder := json.NewDecoder(f)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&meResponse)
	
	if err != nil {
		t.Error(err)
	}
	f.Close()
}
