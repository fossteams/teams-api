package csa

import (
	"encoding/json"
	"fmt"
	api "github.com/fossteams/teams-api/pkg"
	"net/http"
	"net/url"
)

type ChatMessageType string

const (
	ChatMessageTypeMessage ChatMessageType = "Message"
)

type UserEmotion struct {
	Mri   string
	Time  int64 // TODO: Convert to time.time ?
	Value string
}

type Emotion struct {
	Key   string
	Users []UserEmotion
}

type ChatMessageProperties struct {
	Subject               string
	EmailDetails          string
	Meta                  string
	Files                 string
	Emotions              []Emotion
	DeleteTime            int64 // TODO: Convert to time.Time ?
	AdminDelete           bool
	S2SPartnerName        string
	Mentions              string
	Links                 string
	EditTime              interface{} // Can be either string or int64, wtf? TODO: Convert to time.Time ?
	CounterPartyMessageId int64
	OriginContextId       int64
	ParentMessageId       int64
	SkipFanOutToBots      interface{} // Can be either string or bool, wtf?
	Cards                 string
	Importance            string
}

type AnnotationsSummary struct {
	Emotions map[string]int64
}

type ChatMessage struct {
	Id                  string
	SequenceId          int64
	SkypeEditedId       string
	SkypeEditOffset     int
	ClientMessageId     string
	Version             string
	ConversationId      string
	ConversationLink    string
	Type                ChatMessageType
	MessageType         string
	ContentType         string
	Content             string
	AmsReferences       []string
	From                string
	ImDisplayName       string
	S2SPartnerName      string
	ComposeTime         string // TODO: Parse as time.Time
	OriginalArrivalTime string // TODO: Parse as time.Time
	Properties          ChatMessageProperties
	AnnotationsSummary  AnnotationsSummary
}

type MessagesMetadata struct {
	BackwardLink                 string
	SyncState                    string
	LastCompleteSegmentStartTime int64 // TODO: Parse as time.Time
	LastCompleteSegmentEndTime   int64 // TODO: Parse as time.Time
}

type MessagesResponse struct {
	Messages []ChatMessage    `json:"messages"`
	Metadata MessagesMetadata `json:"_metadata"`
}

func (c *CSASvc) GetMessagesByChannel(channel *Channel) ([]ChatMessage, error) {
	endpointUrl := c.getEndpoint(EndpointMessages, "/users/ME/conversations/"+url.QueryEscape(channel.Id)+"/messages")
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

	if resp.StatusCode != http.StatusOK {
		return nil, api.InvalidResponseError(resp)
	}

	jsonBuffer, err := c.getJSON(resp)
	if err != nil {
		return nil, err
	}

	var msgResponse MessagesResponse
	dec := json.NewDecoder(jsonBuffer)
	dec.DisallowUnknownFields()
	err = dec.Decode(&msgResponse)
	if err != nil {
		return nil, fmt.Errorf("unable to decode json: %v", err)
	}

	return msgResponse.Messages, err
}
