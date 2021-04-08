package csa

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type ChannelId string

type PinnedChannelsResponse struct {
	OrderVersion int
	PinChannelOrder []ChannelId
}

func (c *CSASvc) GetPinnedChannels() ([]ChannelId, error) {
	endpointUrl := c.getEndpoint(EndpointChatSvcAgg, "/teams/users/me/pinnedChannels")
	jsonBuffer, err := c.authenticatedGetRequest(endpointUrl)
	if err != nil {
		return nil, err
	}

	var pinnedChannelResponse PinnedChannelsResponse

	d := json.NewDecoder(jsonBuffer)
	d.DisallowUnknownFields()
	err = d.Decode(&pinnedChannelResponse)
	if err != nil {
		return nil, fmt.Errorf("unable to decode JSON: %v", err)
	}
	return pinnedChannelResponse.PinChannelOrder, nil
}


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

var _ sort.Interface = ChannelsByName{}
