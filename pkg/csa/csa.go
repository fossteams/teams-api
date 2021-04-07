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
	token     *api.TeamsToken
	csaSvcUrl *url.URL
	client    *http.Client
	debugSave bool
}

const ChatSvcAgg = "https://teams.microsoft.com/api/csa/"
// Requires an aud:https://chatsvcagg.teams.microsoft.com token

func NewCSAService(token *api.TeamsToken) (*CSASvc, error) {
	// https://teams.microsoft.com/api/csa/api/v1/teams/users/me?isPrefetch=false&enableMembershipSummary=true
	svcUrl, err := url.Parse(ChatSvcAgg)
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, fmt.Errorf("token is nil")
	}

	client := http.DefaultClient

	return &CSASvc{
		csaSvcUrl: svcUrl,
		token:     token,
		client:    client,
	}, nil
}

func (c *CSASvc) DebugSave(debugFlag bool){
	c.debugSave = debugFlag
}

func (c *CSASvc) getEndpoint(path string) *url.URL {
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	endpointUrl, err := c.csaSvcUrl.Parse("api/v1/" + path)
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

	req.Header.Add("Authorization", api.AuthString(c.token))
	return req, nil
}
