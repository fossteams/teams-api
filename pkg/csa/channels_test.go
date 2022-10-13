package csa_test

import (
	"fmt"
	"github.com/fossteams/teams-api/pkg/csa/models"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func DebugSave() bool {
	return os.Getenv("TEAMS_DEBUG_SAVE") == "1"
}

func TestGetMessagesByChannel(t *testing.T) {
	csaSvc := models.initTest(t)
	csaSvc.DebugSave(DebugSave())

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

func TestGetPinnedChannels(t *testing.T) {
	csaSvc := models.initTest(t)
	csaSvc.DebugSave(DebugSave())

	pinnedChannels, err := csaSvc.GetPinnedChannels()
	assert.Nil(t, err)
	assert.NotNil(t, pinnedChannels)

	fmt.Printf("Pinned channels: \n")
	for _, v := range pinnedChannels {
		fmt.Printf("\t%s\n", v)
	}
	fmt.Printf("\n")
}
