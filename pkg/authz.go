package api

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"net/http"
	"reflect"
)

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
	AmdS2S                              string `json:"amdS2S"`
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
	ChatServiceAggregator               string `json:"chatServiceAggregator"`
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
	IsFreemium               bool
	IsBasicLiveEventsEnabled bool
	IsTrial                  bool
	IsAdvComms               bool
}

type AuthzResponse struct {
	Tokens         Tokens         `json:"tokens"`
	Region         Region         `json:"region"`
	Partition      Partition      `json:"partition"`
	RegionGtms     RegionGTMs     `json:"regionGtms"`
	RegionSettings RegionSettings `json:"regionSettings"`
	LicenseDetails LicenseDetails
}

type AuthClient struct {
	client *http.Client
}

func New(client *http.Client) AuthClient {
	if client == nil {
		client = http.DefaultClient
	}

	return AuthClient{client: client}
}

type AuthzType = string

const (
	Refresh AuthzType = "TokenRefresh"
)

func (a *AuthClient) Authz(token *RootSkypeToken, authzType AuthzType) (*SkypeToken, error) {
	req, err := http.NewRequest("POST", TEAMS_API_ENDPOINT+ "/authsvc/v1.0/authz", nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create Authz request: %v", err)
	}

	req.Header.Add("ms-teams-authz-type", authzType)
	req.Header.Add("Authorization", "Bearer " + token.Inner.Raw)
	resp, err := a.client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("unable to perform authz request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		bodyBytes, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return nil, fmt.Errorf("invalid status code returned: 200 expected but %d returned", resp.StatusCode)
		}
		return nil, fmt.Errorf("invalid status code returned: 200 expected but %d returned: %v",
			resp.StatusCode,
			string(bodyBytes),
		)
	}

	dec := json.NewDecoder(resp.Body)
	var authzResp AuthzResponse
	err = dec.Decode(&authzResp)

	if err != nil {
		return nil, fmt.Errorf("unable to decode authz response JSON: %v", err)
	}

	skypeJwt, err := jwt.Parse(authzResp.Tokens.SkypeToken, nil)
	if err != nil {
		shouldThrow := true
		switch err.(type) {
		case *jwt.ValidationError:
			if err.(*jwt.ValidationError).Errors == jwt.ValidationErrorUnverifiable {
				shouldThrow = false
			}
		default:
			fmt.Printf("type=%v", reflect.TypeOf(err))
		}
		if shouldThrow {
			return nil, fmt.Errorf("unable to decode Skype JWT: %v", err)
		}
	}
	skypeToken := SkypeToken{
		Inner: skypeJwt,
		Type:  TokenBearer,
	}
	return &skypeToken, nil
}
