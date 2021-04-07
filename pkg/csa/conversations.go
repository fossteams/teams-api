package csa

import (
	"bytes"
	"encoding/json"
	"fmt"
	api "github.com/fossteams/teams-api/pkg"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
)

func (c *CSASvc) GetConversations() (*ConversationResponse, error) {
	endpointUrl := c.getEndpoint(EndpointChatSvcAgg, "/teams/users/me")

	values := endpointUrl.Query()
	values.Add("isPrefetch", "false")
	values.Add("enableMembershipSummary", "true")
	endpointUrl.RawQuery = values.Encode()

	req, err := c.AuthenticatedRequest("GET", endpointUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, api.InvalidResponseError(resp)
	}

	jsonBuffer, err := c.getJSON(resp)
	if err != nil {
		return nil, err
	}

	var teams ConversationResponse
	decoder := json.NewDecoder(jsonBuffer)
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&teams)

	if err != nil {
		return nil, err
	}

	return &teams, nil
}

func (c *CSASvc) getJSON(resp *http.Response) (io.Reader, error) {
	var jsonBuffer io.Reader

	if c.debugSave {
		// Temporary save response
		f, err := ioutil.TempFile(os.TempDir(), "teams-*.json")
		if err != nil {
			return nil, fmt.Errorf("unable to create temporary file")
		}
		jsonBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("unable to read response body: %v", err)
		}

		_, _ = f.Write(jsonBytes)
		fmt.Printf("saved temporary json to %v\n", f.Name())
		jsonBuffer = bytes.NewReader(jsonBytes)
	} else {
		jsonBuffer = resp.Body
	}
	return jsonBuffer, nil
}


type TeamsByName []Team
type ChannelsByName []Channel

func (c ChannelsByName) Len() int {
	return len(c)
}

func (c ChannelsByName) Less(i, j int) bool {
	return strings.ToLower(c[i].DisplayName) < strings.ToLower(c[j].DisplayName)
}

func (c ChannelsByName) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (t TeamsByName) Len() int {
	return len(t)
}

func (t TeamsByName) Less(i, j int) bool {
	return strings.ToLower(t[i].DisplayName) < strings.ToLower(t[j].DisplayName)
}

func (t TeamsByName) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

var _ sort.Interface = TeamsByName{}
var _ sort.Interface = ChannelsByName{}
