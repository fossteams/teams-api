package mt_test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/fossteams/teams-api/pkg"
	"github.com/fossteams/teams-api/pkg/mt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
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
	email, err := getTokenEMail(t)
	user, err := userSvc.GetUser(email)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	fmt.Printf("user=%#v", user)
	assert.Equal(t, email, user.Email)
}

func getTokenEMail(t *testing.T) (string, error) {
	rootToken, err := api.GetRootToken()
	if err != nil {
		t.Fatalf("unable to get root token: %v", err)
	}

	email := rootToken.Inner.Claims.(jwt.MapClaims)["upn"].(string)
	return email, err
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

func TestGetUserProfilePicture(t *testing.T){
	userSvc := initTest(t)
	email, err := getTokenEMail(t)

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

