package teams_api

import (
	"fmt"
	"github.com/fossteams/teams-api/pkg/csa"
	"github.com/fossteams/teams-api/pkg/models"
	"github.com/fossteams/teams-api/pkg/mt"
	"net/http"
)
import "github.com/fossteams/teams-api/pkg"

type TeamsClient struct {
	httpClient *http.Client
	chatSvc    *csa.CSASvc
	mtSvc      *mt.MTService
}

func (t *TeamsClient) Debug(debugFlag bool) {
	t.chatSvc.DebugSave(debugFlag)
}

func (t *TeamsClient) ChatSvc() *csa.CSASvc {
	return t.chatSvc
}

func New() (*TeamsClient, error) {
	teamsClient := TeamsClient{}

	// Get Skype Spaces Token
	skypeSpaces, err := api.GetSkypeSpacesToken()
	if err != nil {
		return nil, fmt.Errorf("unable to get SkypeSpaces token: %v", err)
	}

	chatSvcToken, err := api.GetChatSvcAggToken()
	if err != nil {
		return nil, fmt.Errorf("unable to get chat service token: %v", err)
	}

	skypeToken, err := api.GetSkypeToken()
	if err != nil {
		return nil, err
	}

	// Initialize Chat Service
	csaSvc, err := csa.NewCSAService(chatSvcToken, skypeToken)
	if err != nil {
		return nil, fmt.Errorf("unable to init Chat Service")
	}
	teamsClient.chatSvc = csaSvc

	// Initialize MT Service
	mtSvc, err := mt.NewMiddleTierService(api.Emea, skypeSpaces)
	if err != nil {
		return nil, fmt.Errorf("unable to init MT Service: %v", err)
	}
	teamsClient.mtSvc = mtSvc

	return &teamsClient, err
}

func (t TeamsClient) GetConversations() (*csa.ConversationResponse, error) {
	return t.chatSvc.GetConversations()
}

func (t TeamsClient) GetMessages(channel *csa.Channel) ([]csa.ChatMessage, error) {
	return t.chatSvc.GetMessagesByChannel(channel)
}

func (t TeamsClient) GetMe() (*models.User, error) {
	return t.mtSvc.GetMe()
}

func (t TeamsClient) FetchShortProfile(mris []string) ([]models.User, error) {
	return t.mtSvc.FetchShortProfile(mris...)
}

func (t TeamsClient) GetTenants() ([]models.Tenant, error) {
	return t.mtSvc.GetTenants()
}

func (t *TeamsClient) GetPinnedChannels() ([]csa.ChannelId, error) {
	return t.chatSvc.GetPinnedChannels()
}
