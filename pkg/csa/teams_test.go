package csa_test

import (
	"fmt"
	"github.com/fossteams/teams-api/pkg"
	"github.com/fossteams/teams-api/pkg/csa"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initTest(t *testing.T) *csa.CSASvc {
	token, err := api.GetChatSvcAggToken()
	if err != nil {
		t.Fatal(err)
	}

	// Get Skype Token
	skypeToken, err := api.GetSkypeToken()
	if err != nil {
		switch err.(type) {
		case api.AuthzError:
			authzError := err.(api.AuthzError).ErrorCode
			if authzError == api.GuestUserNotRedeemed {
				// Auto selecting tenant
				t.Fatalf("please use a token that has already selected one of your tenants")
			}
		}
		t.Errorf("unable to get Skype Token: %v", err)
	}
	skypeToken.Type = api.TokenSkype

	csaSvc, err := csa.NewCSAService(token, skypeToken)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
	return csaSvc
}

func TestGetConversations(t *testing.T) {
	csaSvc := initTest(t)
	csaSvc.DebugSave(true)

	conversations, err := csaSvc.GetConversations()

	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, conversations)
	assert.Greater(t, len(conversations.Chats), 0)
	assert.Greater(t, len(conversations.Teams), 0)
	fmt.Printf("Conversations: %+v\n", conversations)
}
