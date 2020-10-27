package auth

type Tokens struct {
	SkypeToken string `json:"skypeToken"`
	ExpiresIn  int    `json:"expiresIn"`
}

type Region string
type Partition string

const (
	Emea Region = "emea"
)

const (
	Emea1 Partition = "emea01"
)

type RegionGTMs struct {
	Ams                                 string `json:"ams"`
	AmsV2                               string `json:"amsV2"`
	AmdS2S                               string `json:"amdS2S"`
	AmsS2S                              string `json:"amsS2S"`
	AppsDataLayerService                string `json:"appsDataLayerService"`
	AppsDataLayerServiceS2S             string `json:"appsDataLayerServiceS2S"`
	CallingCallControllerServiceURL     string `json:"calling_callControllerServiceUrl"`
	CallingCallStoreUrl                 string `json:"calling_callStoreUrl"`
	CallingConversationServiceURL       string `json:"calling_conversationServiceUrl"`
	CallingKeyDistributionUrl           string `json:"calling_keyDistributionUrl"`
	CallingPotentialCallRequestUrl      string `json:"calling_potentialCallRequestUrl"`
	CallingSharedLineOptionsUrl         string `json:"calling_sharedLineOptionsUrl"`
	CallingUdpTransportUrl              string `json:"calling_udpTransportUrl"`
	CallingUploadLogRequestUrl          string `json:"calling_uploadLogRequestUrl"`
	CallingS2SBroker                    string `json:"callingS2S_Broker"`
	CallingS2SCallController            string `json:"callingS2S_CallController"`
	CallingS2SCallStore                 string `json:"callingS2S_CallStore"`
	CallingS2SContentSharing            string `json:"callingS2S_ContentSharing"`
	CallingS2SConversationService       string `json:"callingS2S_ConversationService"`
	CallingS2SEnterpriseProxy           string `json:"callingS2S_EnterpriseProxy"`
	CallingS2SMediaController           string `json:"callingS2S_MediaController"`
	CallingS2SPlatformMediaAgent        string `json:"callingS2S_PlatformMediaAgent"`
	ChatService                         string `json:"chatService"`
	ChatServiceAggregator				string `json:"chatServiceAggregator"`
	ChatServiceS2S                      string `json:"chatServiceS2S"`
	Drad                                string `json:"drad"`
	MailHookS2S                         string `json:"mailhookS2S"`
	MiddleTier                          string `json:"middleTier"`
	MiddleTierS2S                       string `json:"middleTierS2S"`
	MtImageService                      string `json:"mtImageService"`
	PowerPointStateService              string `json:"powerPointStateService"`
	Search                              string `json:"search"`
	SearchTelemetry                     string `json:"searchTelemetry"`
	TeamsAndChannelsService             string `json:"teamsAndChannelsService"`
	TeamsAndChannelsProvisioningService string `json:"teamsAndChannelsProvisioningService"`
	Urlp                                string `json:"urlp"`
	UrlpV2                              string `json:"urlpV2"`
	UnifiedPresence                     string `json:"unifiedPresence"`
	UserEntitlementService              string `json:"userEntitlementService"`
	UserIntelligenceService             string `json:"userIntelligenceService"`
	UserProfileService                  string `json:"userProfileService"`
	UserProfileServiceS2S               string `json:"userProfileServiceS2S"`
}

type RegionSettings struct {
	IsUnifiedPresenceEnabled        bool
	IsOutOfOfficeIntegrationEnabled bool
	IsContactMigrationEnabled       bool
	IsAppsDiscoveryEnabled          bool
	IsFederationEnabled             bool
}

type LicenseDetails struct {
	IsFreemium bool
	IsBasicLiveEventsEnabled bool
	IsTrial bool
	IsAdvComms bool
}

type AuthzResponse struct {
	Tokens     Tokens     `json:"tokens"`
	Region     Region     `json:"region"`
	Partition  Partition  `json:"partition"`
	RegionGtms RegionGTMs `json:"regionGtms"`
	RegionSettings RegionSettings `json:"regionSettings"`
	LicenseDetails LicenseDetails
}
