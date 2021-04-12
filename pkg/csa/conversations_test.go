package csa_test

import (
	"encoding/json"
	"fmt"
	"github.com/fossteams/teams-api/pkg/csa"
	"os"
	"testing"
)

func TestParseConversations(t *testing.T) {
	f, err := os.Open("../../resources/chatsvcagg/conversations/conversations-1.json")
	defer f.Close()
	if err != nil {
		t.Fatalf("unable to open file: %v", err)
	}

	var conversations csa.ConversationResponse
	dec := json.NewDecoder(f)

	err = dec.Decode(&conversations)
	dec.DisallowUnknownFields()
	if err != nil {
		t.Fatalf("unable to decode JSON: %v", err)
	}
	fmt.Printf("conversations:\n%+v\n", conversations)
}