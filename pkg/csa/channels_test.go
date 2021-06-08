package csa_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMessagesByChannel(t *testing.T) {
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

func TestGetPinnedChannels(t *testing.T) {
	csaSvc := initTest(t)
	csaSvc.DebugSave(true)
	pinnedChannels, err := csaSvc.GetPinnedChannels()
	assert.Nil(t, err)
	assert.NotNil(t, pinnedChannels)

	fmt.Printf("Pinned channels: \n")
	for _, v := range pinnedChannels {
		fmt.Printf("\t%s", v)
	}
	fmt.Printf("\n")
}
