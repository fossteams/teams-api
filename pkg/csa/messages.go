package csa

import (
	"encoding/json"
	"fmt"
	"github.com/fossteams/teams-api/pkg/errors"
	models2 "github.com/fossteams/teams-api/pkg/models"
	"github.com/fossteams/teams-api/pkg/util"
	"net/http"
	"net/url"
)

func (c *CSASvc) GetMessagesByChannel(channel *models2.Channel) ([]models2.ChatMessage, error) {
	endpointUrl := c.getEndpoint(EndpointMessages,
		fmt.Sprintf("/users/ME/conversations/%s/messages",
			url.PathEscape(channel.Id),
		),
	)
	values := endpointUrl.Query()
	values.Add("view", "msnp24Equivalent|supportsMessageProperties")
	values.Add("pageSize", "200")
	values.Add("startTime", "1")
	endpointUrl.RawQuery = values.Encode()

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

	jsonBuffer, err := util.GetJSON(resp, c.debugSave)
	if err != nil {
		return nil, err
	}

	var msgResponse models2.MessagesResponse
	dec := json.NewDecoder(jsonBuffer)
	if c.debugDisallowUnknownFields {
		dec.DisallowUnknownFields()
	}
	err = dec.Decode(&msgResponse)
	if err != nil {
		return nil, fmt.Errorf("unable to decode json: %v", err)
	}

	return msgResponse.Messages, err
}
