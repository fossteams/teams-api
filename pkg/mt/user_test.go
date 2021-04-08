package mt_test

import (
	"fmt"
	"github.com/fossteams/teams-api/pkg"
	"github.com/fossteams/teams-api/pkg/mt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func initTest(t *testing.T) *mt.MTService {
	token, err := api.GetRootToken()
	if err != nil {
		t.Error(err)
	}
	
	userSvc, err := mt.NewMiddleTierService(api.Emea, token)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
	return userSvc
}

func TestGetUser(t *testing.T){
	userSvc := initTest(t)

	userEmail := os.Getenv("MS_TEAMS_USER_EMAIL")
	user, err := userSvc.GetUser(userEmail)
	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestGetMe(t *testing.T){
	userSvc := initTest(t)

	user, err := userSvc.GetMe()
	assert.Nil(t, err)
	assert.NotNil(t, user)
	fmt.Printf("user=%#v\n", user)
}

func TestFetchShortProfiles(t *testing.T){
	userSvc := initTest(t)

	mrisEnv := os.Getenv("MS_TEAMS_MRIS")
	mris := strings.Split(mrisEnv, ",")
	
	user, err := userSvc.FetchShortProfile(mris...)
	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestGetUserProfilePicture(t *testing.T){
	userSvc := initTest(t)

	userEmail := os.Getenv("MS_TEAMS_USER_EMAIL")
	profilePicture, err := userSvc.GetProfilePicture(userEmail)
	assert.Nil(t, err)
	assert.NotNil(t, profilePicture)
	assert.Greater(t, len(profilePicture), 0)
	f, err := ioutil.TempFile(os.TempDir(), "teams-pkg*.jpg")
	if err != nil {
		t.Errorf("unable to create temp file: %v", err)
		t.Fail()
	}
	
	_, _ = f.Write(profilePicture)
	_ = f.Close()
	fmt.Printf("profile picture saved to %v", f.Name())
}

