package teams_api

import (
	"fmt"
	"github.com/fossteams/teams-api/pkg/csa"
	"net/http"
)
import "github.com/fossteams/teams-api/pkg"

type TeamsClient struct {
	httpClient *http.Client
	chatSvc *csa.CSASvc
}

func (c *TeamsClient) Debug(debugFlag bool) {
	c.chatSvc.DebugSave(debugFlag)
}

func New() (*TeamsClient, error) {
	teamsClient := TeamsClient{}

	// Get Teams Token
	_, err := api.GetTeamsToken()
	if err != nil {
		return nil, fmt.Errorf("unable to get teams token: %v", err)
	}

	chatSvcToken, err := api.GetChatSvcAggToken()
	if err != nil {
		return nil, fmt.Errorf("unable to get chat service token: %v", err)
	}

	skypeToken, err := api.GetSkypeToken()
	if err != nil {
		return nil, fmt.Errorf("unable to get SkypeToken: %v", err)
	}

	// Initialize Chat Service
	csaSvc, err := csa.NewCSAService(chatSvcToken, skypeToken)
	if err != nil {
		return nil, fmt.Errorf("unable to init Chat Service")
	}

	teamsClient.chatSvc = csaSvc
	return &teamsClient, err
}

func (t *TeamsClient) GetConversations() (*csa.ConversationResponse, error) {
	return t.chatSvc.GetConversations()
}

func (t *TeamsClient) GetMessages(channel *csa.Channel) ([]csa.ChatMessage, error) {
	return t.chatSvc.GetMessagesByChannel(channel)
}