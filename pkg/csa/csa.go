package csa

import (
	"fmt"
	"github.com/fossteams/teams-api/pkg"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type CSASvc struct {
	token                      *api.TeamsToken
	csaSvcUrl                  *url.URL
	msgUrl                     *url.URL
	client                     *http.Client
	debugSave                  bool
	skypeToken                 *api.TeamsToken
	debugDisallowUnknownFields bool
}

const ChatSvcAgg = "https://teams.microsoft.com/api/csa/api/"
const MessagesHost = "https://emea.ng.msg.teams.microsoft.com/"

// Requires an aud:https://chatsvcagg.teams.microsoft.com token

func NewCSAService(token *api.TeamsToken, skypeToken *api.SkypeToken) (*CSASvc, error) {
	// https://teams.microsoft.com/api/csa/api/v1/teams/users/me?isPrefetch=false&enableMembershipSummary=true
	svcUrl, err := url.Parse(ChatSvcAgg)
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, fmt.Errorf("token cannot be nil")
	}

	if skypeToken == nil {
		return nil, fmt.Errorf("skypeToken cannot be nil")
	}

	msgUrl, err := url.Parse(MessagesHost)
	if err != nil {
		return nil, fmt.Errorf("unable to parse Messages URL: %v", err)
	}

	client := http.DefaultClient

	return &CSASvc{
		csaSvcUrl:                  svcUrl,
		msgUrl:                     msgUrl,
		token:                      token,
		skypeToken:                 skypeToken,
		client:                     client,
		debugDisallowUnknownFields: false,
	}, nil
}

func (c *CSASvc) DebugSave(debugFlag bool) {
	c.debugSave = debugFlag
}

func (c *CSASvc) DebugDisallowUnknownFields(debugFlag bool) {
	c.debugDisallowUnknownFields = debugFlag
}

type EndpointType string

const (
	EndpointChatSvcAgg EndpointType = "chatsvcagg"
	EndpointMessages   EndpointType = "messages"
)

func (c *CSASvc) getEndpoint(t EndpointType, path string) *url.URL {
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	var url = c.csaSvcUrl
	switch t {
	case EndpointChatSvcAgg:
		url = c.csaSvcUrl
	case EndpointMessages:
		url = c.msgUrl
	}
	endpointUrl, err := url.Parse("v1/" + path)
	if err != nil {
		return c.csaSvcUrl
	}

	return endpointUrl
}

func (c *CSASvc) AuthenticatedRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(url, ChatSvcAgg) {
		// Use ChatSvgAgg Token
		req.Header.Add("Authorization", api.AuthString(c.token))
	} else if strings.HasPrefix(url, MessagesHost) {
		// Use SkypeToken
		req.Header.Add("Authentication", api.AuthString(c.skypeToken))
	}

	return req, nil
}
