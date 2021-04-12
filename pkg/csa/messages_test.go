package csa_test

import (
	"encoding/json"
	"fmt"
	"github.com/fossteams/teams-api/pkg/csa"
	"os"
	"testing"
)

func TestParseMessages(t *testing.T) {
	f, err := os.Open("../../resources/chatsvcagg/messages/messages-1.json")
	defer f.Close()
	if err != nil {
		t.Fatalf("unable to open file: %v", err)
	}

	var messages csa.MessagesResponse
	dec := json.NewDecoder(f)

	err = dec.Decode(&messages)
	dec.DisallowUnknownFields()
	if err != nil {
		t.Fatalf("unable to decode JSON: %v", err)
	}
	fmt.Printf("messages:\n%+v\n", messages)
}