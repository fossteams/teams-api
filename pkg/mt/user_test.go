package mt_test

import (
	"encoding/json"
	"fmt"
	"github.com/fossteams/teams-api/pkg"
	"github.com/fossteams/teams-api/pkg/models"
	"github.com/fossteams/teams-api/pkg/mt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func initTest(t *testing.T) *mt.Service {
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

func TestGetUser(t *testing.T) {
	userSvc := initTest(t)
	userSvc.DebugDisallowUnknownFields(true)
	userSvc.DebugSave(true)
	email, err := getTokenEmail(t)
	user, err := userSvc.GetUser(email)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	fmt.Printf("user=%#v", user)
	assert.Equal(t, email, user.Email)
}

func TestParseUsersResponse(t *testing.T) {
	f, err := os.Open("../../resources/mt/user/user-1.json")
	defer f.Close()
	if err != nil {
		t.Fatalf("unable to open file: %v", err)
	}

	var typedEntry = struct {
		Value models.User
		Type  string
	}{}

	var user models.User
	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()
	err = dec.Decode(&typedEntry)
	if err != nil {
		t.Fatalf("unable to decode JSON: %v", err)
	}

	assert.Equal(t, "Microsoft.SkypeSpaces.MiddleTier.Models.AadMember", typedEntry.Type)
	user = typedEntry.Value

	fmt.Printf("user:%+v\n", user)
	assert.NotNil(t, user)
	assert.Equal(t, "Denys", user.GivenName)
	assert.Equal(t, "Vitali", user.Surname)
	assert.Equal(t, "Denys Vitali", user.DisplayName)
	assert.Equal(t, "teams-cli@outlook.com", user.Email)
	assert.True(t, user.SkypeTeamsInfo.IsSkypeTeamsUser)
	assert.True(t, user.AccountEnabled)
	assert.True(t, user.IsSipDisabled)
	assert.False(t, user.IsShortProfile)
	assert.Equal(t, "8:orgid:fa814989-41d0-4d4b-a365-e5f44e406847", user.Mri)
}

func getTokenEmail(t *testing.T) (string, error) {
	rootToken, err := api.GetRootToken()
	if err != nil {
		t.Fatalf("unable to get root token: %v", err)
	}

	return mt.GetTokenEmail(rootToken)
}

func TestGetMe(t *testing.T) {
	userSvc := initTest(t)

	user, err := userSvc.GetMe()
	assert.Nil(t, err)
	assert.NotNil(t, user)
	fmt.Printf("user=%#v\n", user)
}

func TestFetchShortProfiles(t *testing.T) {
	userSvc := initTest(t)
	user, err := userSvc.GetMe()
	if err != nil {
		t.Fatalf("unable to get me: %v", err)
	}

	mris := []string{user.Mri}

	users, err := userSvc.FetchShortProfile(mris...)
	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, 1, len(users))

	assert.Equal(t, user.Email, users[0].Email)
	assert.Equal(t, user.DisplayName, users[0].DisplayName)

	fmt.Printf("users=%#v\n", users)

}

func TestGetUserProfilePicture(t *testing.T) {
	userSvc := initTest(t)
	email, err := getTokenEmail(t)

	profilePicture, err := userSvc.GetProfilePicture(email)
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
