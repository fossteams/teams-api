package models

import (
	"sort"
	"time"
)

type ChatMessageType string

const (
	ChatMessageTypeMessage  ChatMessageType = "Message"
	EvenCall                ChatMessageType = "Event/Call"
	ThreadActivityAddMember ChatMessageType = "ThreadActivity/AddMember"
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

type DeltaEmotion struct {
	Key   string
	Users []UserEmotion
}

type ChatMessageProperties struct {
	Subject               string
	Title                 string
	EmailDetails          string
	Meta                  string
	Files                 string
	Emotions              []Emotion
	DeltaEmotions         []DeltaEmotion
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
	Atp                   string
	CrossPostId           string
	Meeting               string
	SkypeGuid             string
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
	FromTenantId        string
	ImDisplayName       string
	S2SPartnerName      string
	ComposeTime         RFC3339Time
	OriginalArrivalTime RFC3339Time
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

type SortMessageByTime []ChatMessage

func (s SortMessageByTime) Len() int {
	return len(s)
}

func (s SortMessageByTime) Less(i, j int) bool {
	ti := time.Time(s[i].ComposeTime)
	tj := time.Time(s[j].ComposeTime)
	return ti.Before(tj)
}

func (s SortMessageByTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var _ sort.Interface = SortMessageByTime{}
