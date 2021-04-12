package csa

import (
	"encoding/json"
	"fmt"
	"github.com/fossteams/teams-api/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func (c *CSASvc) GetConversations() (*ConversationResponse, error) {
	endpointUrl := c.getEndpoint(EndpointChatSvcAgg, "/teams/users/me")

	values := endpointUrl.Query()
	values.Add("isPrefetch", "false")
	values.Add("enableMembershipSummary", "true")
	endpointUrl.RawQuery = values.Encode()

	jsonBuffer, err := c.authenticatedGetRequest(endpointUrl)
	if err != nil {
		return nil, err
	}

	var teams ConversationResponse
	decoder := json.NewDecoder(jsonBuffer)
	if c.debugDisallowUnknownFields {
		decoder.DisallowUnknownFields()
	}
	err = decoder.Decode(&teams)

	if err != nil {
		return nil, err
	}

	return &teams, nil
}

func (c *CSASvc) authenticatedGetRequest(endpointUrl *url.URL) (io.Reader, error) {
	req, err := c.AuthenticatedRequest("GET", endpointUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		return nil, errors.NewHTTPError(expectedStatusCode, resp.StatusCode, nil)
	}

	jsonBuffer, err := c.getJSON(resp)
	if err != nil {
		return nil, fmt.Errorf("unable to read JSON: %v", err)
	}

	return jsonBuffer, nil
}

type TeamsByName []Team

func (t TeamsByName) Less(i, j int) bool {
	return strings.ToLower(t[i].DisplayName) < strings.ToLower(t[j].DisplayName)
}

func (t TeamsByName) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

var _ sort.Interface = TeamsByName{}
