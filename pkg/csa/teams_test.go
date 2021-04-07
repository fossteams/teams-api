package csa

import (
	"encoding/json"
	"fmt"
	"github.com/fossteams/teams-api/pkg"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)
func initTest(t *testing.T) *CSASvc {
	token, err := api.GetChatSvcAggToken()
	if err != nil {
		t.Error(err)
	}

	// Get Skype Token
	skypeToken, err := api.GetSkypeToken()
	if err != nil {
		t.Errorf("unable to get Skype Token: %v", err)
	}
	skypeToken.Type = api.TokenSkype

	csaSvc, err := NewCSAService(token, skypeToken)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
	return csaSvc
}

func TestGetConversations(t *testing.T){
	csaSvc := initTest(t)
	csaSvc.DebugSave(true)

	conversations, err := csaSvc.GetConversations()

	if err != nil {
		t.Error(err)
	}


	assert.NotNil(t, conversations)
	assert.Greater(t, len(conversations.Chats), 0)
	assert.Greater(t, len(conversations.Teams), 0)
	fmt.Printf("Conversations: %+v", conversations)
}

func TestGetMessages(t *testing.T){
	csaSvc := initTest(t)
	csaSvc.DebugSave(true)

	conversations, err := csaSvc.GetConversations()
	if err != nil {
		t.Fatalf("unable to get conversations: %v", err)
	}


	assert.NotNil(t, conversations)
	assert.Greater(t, len(conversations.Chats), 0)
	assert.Greater(t, len(conversations.Teams), 0)

	messages, err := csaSvc.GetMessagesByChannel(&conversations.Teams[0].Channels[0])
	if err != nil {
		t.Fatalf("unable to get messages by channel: %v", err)
	}

	assert.Greater(t, len(messages), 0)
}

func TestParseConversations(t *testing.T) {
	f, err := os.Open("../../resources/chatsvcagg/conversations/conversations-1.json")
	if err != nil {
		t.Fatalf("unable to open file: %v", err)
	}

	var conversations ConversationResponse
	dec := json.NewDecoder(f)
	//dec.DisallowUnknownFields()

	err = dec.Decode(&conversations)
	if err != nil {
		t.Fatalf("unable to decode JSON: %v", err)
	}
	fmt.Printf("conversations:\n%+v\n", conversations)
}

func TestDecodeTeamsJson(t *testing.T) {
	f, err := os.Open("../../resources/chatsvcagg/conversations/conversations-1.json")
	if err != nil {
		t.Error(err)
	}
	var meResponse ConversationResponse
	decoder := json.NewDecoder(f)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&meResponse)
	
	if err != nil {
		t.Error(err)
	}
	f.Close()
}
