package csa

import (
	"encoding/json"
	"github.com/fossteams/teams-api/api"
	"net/http"
	"time"
)

type User struct {
}

type FeedProperty struct {
	IsEmptyConversation        string
	ConsumptionHorizon         string
	ConsumptionHorizonBookmark string
}

type PrivateFeed struct {
	Id          string
	Type        string
	Version     int
	Properties  FeedProperty
	LastMessage Message
	Messages    string
	TargetLink  string
	StreamType  string
}

type ConversationMetadata struct {
	SyncToken string
}

type ChatMemberRole string

const (
	ChatMemberAdmin ChatMemberRole = "Admin"
)

type ChatMember struct {
	IsMuted      bool
	Mri          string
	Role         ChatMemberRole
	FriendlyName string
	TenantId     string
}

type DateRange struct {
	StartDate time.Time
	EndDate   time.Time
}

type WeeklyRecurrence struct {
	Interval      int
	DaysOfTheWeek []int
}

type MonthlyRecurrence struct {
	Interval            int
	WeekOfTheMonthIndex int
	DayOfTheWeek        int
}

type RecurrencePattern struct {
	PatternType     int
	Weekly          WeeklyRecurrence
	RelativeMonthly MonthlyRecurrence
}

type MeetingInfo struct {
	Subject                string
	Location               string
	StartTime              string // TODO: Switch to time.Time
	EndTime                string // TODO: Switch to time.Time
	ExchangeId             *string
	ICalUid                string
	IsCancelled            bool
	EventRecurrenceRange   DateRange
	EventRecurrencePattern RecurrencePattern
	AppointmentType        int
	MeetingType            int
	Scenario               string
}

type Chat struct {
	ChatSubType               int                `json:"chatSubType"`
	ChatType                  string             `json:"chatType"`
	ConsumptionHorizon        ConsumptionHorizon `json:"consumptionHorizon"`
	ConversationBlockedAt     int                `json:"conversationBlockedAt"`
	CreatedAt                 string             `json:"createdAt"`
	Creator                   string             `json:"creator"`
	HasTranscript             bool               `json:"hasTranscript"`
	Hidden                    bool               `json:"hidden"`
	Id                        string             `json:"id"`
	InteropConversationStatus string             `json:"interopConversationStatus"`
	InteropType               int                `json:"interopType"`
	IsDisabled                bool               `json:"isDisabled"`
	IsGapDetectionEnabled     bool               `json:"isGapDetectionEnabled"`
	IsHighImportance          bool               `json:"isHighImportance"`
	IsLastMessageFromMe       bool               `json:"isLastMessageFromMe"`
	IsMessagingDisabled       bool               `json:"isMessagingDisabled"`
	IsOneOnOne                bool               `json:"isOneOnOne"`
	IsRead                    bool               `json:"isRead"`
	IsSticky                  bool               `json:"isSticky"`
	LastJoinAt                time.Time          `json:"lastJoinAt"`
	LastLeaveAt               time.Time          `json:"lastLeaveAt"`
	LastMessage               Message            `json:"lastMessage"`
	MeetingInformation        MeetingInfo        `json:"meetingInformation"`
	MeetingPolicy             string             `json:"meetingPolicy"`
	Members                   []ChatMember       `json:"members"`
	RetentionHorizon          string             `json:"retentionHorizon"`
	RetentionHorizonV2        string             `json:"retentionHorizonV2"`
	ShareHistoryFromTime      string             `json:"shareHistoryFromTime"`
	Tabs                      []Tab              `json:"tabs"`
	TenantId                  string             `json:"tenantId"`
	ThreadVersion             int                `json:"threadVersion"`
	ThreadSchemaVersion       string             `json:"threadSchemaVersion,omitempty"`
	Title                     string             `json:"title"`
	UserConsumptionHorizon    ConsumptionHorizon `json:"userConsumptionHorizon"`
	Version                   int                `json:"version"`
}

type ConversationResponse struct {
	Chats        []Chat
	Metadata     ConversationMetadata
	PrivateFeeds []PrivateFeed
	Teams        []Team
	Users        []User
}

type ConsumptionHorizon struct {
	OriginalArrivalTime int
	TimeStamp           int
	ClientMessageId     string
}

type RetentionHorizon struct {
}

type RetentionHorizonV2 struct {
}

type FileSettings struct {
	FilesRelativePath string
	DocumentLibraryId string
}

type Tab struct {
	Id           string
	Name         string
	DefinitionId string
	Directive    string
	TabType      string
	Order        float32
	ReplyChainId int
	Settings     interface{}
}

type MessageType string

const (
	TextMessage MessageType = "Text"
)

type Message struct {
	MessageType             MessageType
	Content                 string
	ClientMessageId         string
	ImDisplayName           string
	Id                      string
	Type                    string
	ComposeTime             time.Time
	OriginalArrivalTime     time.Time
	ContainerId             string
	ParentMessageId         string
	From                    string
	SequenceId              int
	Version                 int
	ThreadType              *string
	IsEscalationToNewPerson bool
}

type MemberRole int

const (
	Unknown MemberRole = iota
	Member
)

type ChannelType int

const (
	NormalChannel ChannelType = iota
)

type ConnectorProfile struct {
	AvatarUrl     string
	DisplayName   string
	IncomingUrl   *string
	ConnectorType string
	Id            string
}

type ChannelSettings struct {
	ChannelPostPermissions           int
	ChannelReplyPermissions          int
	ChannelPinPostPermissions        int
	ChannelConnectorsPostPermissions int
	ChannelBotsPostPermissions       int
}

type Channel struct {
	Id                       string
	DisplayName              string
	Description              string
	ConsumptionHorizon       *ConsumptionHorizon
	RetentionHorizon         *RetentionHorizon
	RetentionHorizonV2       *RetentionHorizonV2
	Version                  int
	ThreadVersion            int
	ParentTeamId             string
	IsGeneral                bool
	IsFavorite               bool
	IsFollowed               bool
	IsMember                 bool
	Creator                  string
	IsMessageRead            bool
	IsImportantMessageRead   bool
	IsGapDetectionEnabled    bool
	DefaultFileSettings      FileSettings
	Tabs                     []Tab
	LastMessage              Message
	ConnectorProfiles        []ConnectorProfile
	IsDeleted                bool
	IsPinned                 bool
	LastImportantMessageTime time.Time
	LastLeaveAt              time.Time
	LastJoinAt               time.Time
	MemberRole               MemberRole
	IsMuted                  bool
	MembershipExpiry         int
	IsFavoriteByDefault      bool
	CreationTime             time.Time
	IsArchived               bool
	ChannelType              ChannelType
	ChannelSettings          ChannelSettings
	MembershipVersion        int
	MembershipSummary        *string
	IsModerator              bool
	GroupId                  string
	ChannelOnlyMember        bool
	ThreadSchemaVersion      string `json:"threadSchemaVersion,omitempty"`
	UserConsumptionHorizon   ConsumptionHorizon
}

type AccessType int

const (
	NormalAccess AccessType = 3
)

type TeamSettings struct {
	AddDisplayContent               bool `json:"addDisplayContent"`
	AdminDeleteEnabled              bool `json:"adminDeleteEnabled"`
	ChannelMention                  bool `json:"channelMention"`
	CreateIntegration               bool `json:"createIntegration"`
	CreateTab                       bool `json:"createTab"`
	CreateTopic                     bool `json:"createTopic"`
	CustomMemesEnabled              bool `json:"customMemesEnabled"`
	DeleteEnabled                   bool `json:"deleteEnabled"`
	DeleteIntegration               bool `json:"deleteIntegration"`
	DeleteTab                       bool `json:"deleteTab"`
	DeleteTopic                     bool `json:"deleteTopic"`
	EditEnabled                     bool `json:"editEnabled"`
	GeneralChannelPosting           int  `json:"generalChannelPosting"`
	GiphyEnabled                    bool `json:"giphyEnabled"`
	GiphyRating                     int  `json:"giphyRating"`
	InstallApp                      bool `json:"installApp"`
	IsPrivateChannelCreationEnabled bool `json:"isPrivateChannelCreationEnabled"`
	MessageThreadingEnabled         bool `json:"messageThreadingEnabled"`
	RemoveDisplayContent            bool `json:"removeDisplayContent"`
	StickersEnabled                 bool `json:"stickersEnabled"`
	TeamMemesEnabled                bool `json:"teamMemesEnabled"`
	TeamMention                     bool `json:"teamMention"`
	UninstallApp                    bool `json:"uninstallApp"`
	UpdateIntegration               bool `json:"updateIntegration"`
	UpdateTopic                     bool `json:"updateTopic"`
	UploadCustomApp                 bool `json:"uploadCustomApp"`
}

type TeamStatus struct {
	ExchangeTeamCreationStatus   int
	SharePointSiteCreationStatus int
	TeamNotebookCreationStatus   int
	ExchangeTeamDeletionStatus   int
}

type MembershipSummary struct {
	BotCount          int
	MutedMembersCount int
	TotalMemberCount  int
	AdminRoleCount    int
	UserRoleCount     int
	GuestRoleCount    int
}

type TeamSiteInformation struct {
	GroupId              string
	SharepointSiteUrl    string
	NotebookId           string
	IsOneNoteProvisioned bool
}

type TeamType int

const (
	NormalTeam TeamType = 0
)

type ExtensionDefinition struct {
	UpdatedTime time.Time
}

type Team struct {
	AccessType                     AccessType          `json:"accessType"`
	ChannelOnlyMember              bool                `json:"channelOnlyMember"`
	Channels                       []Channel           `json:"channels"`
	Classification                 string              `json:"classification"`
	ConversationVersion            string              `json:"conversationVersion"`
	Creator                        string              `json:"creator"`
	Description                    string              `json:"description"`
	DisplayName                    string              `json:"displayName"`
	DynamicMembership              bool                `json:"dynamicMembership"`
	GuestUsersCategory             string              `json:"guestUsersCategory"`
	Id                             string              `json:"id"`
	IsArchived                     bool                `json:"isArchived"`
	IsCollapsed                    bool                `json:"isCollapsed"`
	IsCreator                      bool                `json:"isCreator"`
	IsDeleted                      bool                `json:"isDeleted"`
	IsFavorite                     bool                `json:"isFavorite"`
	IsFollowed                     bool                `json:"isFollowed"`
	IsTeamLocked                   bool                `json:"isTeamLocked"`
	IsTenantWide                   bool                `json:"isTenantWide"`
	IsUnlockMembershipSyncRequired bool                `json:"isUnlockMembershipSyncRequired"`
	IsUserMuted                    bool                `json:"isUserMuted"`
	LastJoinAt                     string              `json:"lastJoinAt"`
	MaximumMemberLimitExceeded     bool                `json:"maximumMemberLimitExceeded"`
	MemberRole                     MemberRole          `json:"memberRole"`
	MembershipExpiry               int                 `json:"membershipExpiry"`
	MembershipSummary              MembershipSummary   `json:"membershipSummary"`
	MembershipVersion              int                 `json:"membershipVersion"`
	PictureETag                    string              `json:"pictureETag"`
	SmtpAddress                    string              `json:"smtpAddress"`
	TeamGuestSettings              TeamSettings        `json:"teamGuestSettings"`
	TeamSettings                   TeamSettings        `json:"teamSettings"`
	TeamSiteInformation            TeamSiteInformation `json:"teamSiteInformation"`
	TeamStatus                     TeamStatus          `json:"teamStatus"`
	TeamType                       TeamType            `json:"teamType"`
	ExtensionDefinition            ExtensionDefinition `json:"extensionDefinition"`
	TenantId                       string              `json:"tenantId"`
	ThreadSchemaVersion            string              `json:"threadSchemaVersion,omitempty"`
	ThreadVersion                  string              `json:"threadVersion"`
}

func (c *CSASvc) GetConversations() (*ConversationResponse, error) {
	endpointUrl := c.getEndpoint("/teams/users/me")

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

	var teams ConversationResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&teams)

	if err != nil {
		return nil, err
	}

	return &teams, nil
}
