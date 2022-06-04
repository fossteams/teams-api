package models

type UserResponse struct {
	Value User   `json:"value"`
	Type  string `json:"type"`
}

type SkypeTeamsInfo struct {
	IsSkypeTeamsUser bool
}

type FeatureSettings struct {
	IsPrivateChatEnabled           bool
	EnableShiftPresence            bool
	CoExistenceMode                string
	EnableScheduleOwnerPermissions bool
}

type PhoneType string

const (
	MobilePhone PhoneType = "Mobile"
)

type Phone struct {
	Type   PhoneType `json:"type"`
	Number string    `json:"number"`
}

type User struct {
	Alias                      string          `json:"alias"`
	AccountEnabled             bool            `json:"accountEnabled,omitempty"`
	Department                 string          `json:"department,omitempty"`
	DisplayName                string          `json:"displayName"`
	Email                      string          `json:"email"`
	FeatureSettings            FeatureSettings `json:"featureSettings,omitempty"`
	GivenName                  string          `json:"givenName"`
	IsShortProfile             bool            `json:"isShortProfile"`
	IsSipDisabled              bool            `json:"isSipDisabled,omitempty"`
	JobTitle                   string          `json:"jobTitle"`
	Mail                       string          `json:"mail,omitempty"`
	Mobile                     string          `json:"mobile,omitempty"`
	Mri                        string          `json:"mri,omitempty"`
	ObjectId                   string          `json:"objectId"`
	ObjectType                 string          `json:"objectType,omitempty"`
	Phones                     []Phone         `json:"phones,omitempty"`
	PhysicalDeliveryOfficeName string          `json:"physicalDeliveryOfficeName,omitempty"`
	PreferredLanguage          string          `json:"preferredLanguage,omitempty"`
	ResponseSourceInformation  string          `json:"responseSourceInformation,omitempty"`
	ShowInAddressList          bool            `json:"showInAddressList,omitempty"`
	SipProxyAddress            string          `json:"sipProxyAddress,omitempty"`
	SkypeTeamsInfo             SkypeTeamsInfo  `json:"skypeTeamsInfo,omitempty"`
	SmtpAddresses              []string        `json:"smtpAddresses,omitempty"`
	Surname                    string          `json:"surname"`
	TenantName                 string          `json:"tenantName"`
	TelephoneNumber            string          `json:"telephoneNumber,omitempty"`
	Type                       string          `json:"type"`
	UserLocation               string          `json:"userLocation"`
	UserPrincipalName          string          `json:"userPrincipalName"`
	UserType                   string          `json:"userType,omitempty"`
}

type UserType string

const (
	Member UserType = "member"
)

type TenantType string

const (
	Organization TenantType = "organization"
)

type Tenant struct {
	TenantID             string     `json:"tenantId"`
	TenantName           string     `json:"tenantName"`
	UserId               string     `json:"userId"`
	IsInvitationRedeemed bool       `json:"isInvitationRedeemed"`
	CountryLetterCode    string     `json:"countryLetterCode"`
	UserType             UserType   `json:"userType"`
	TenantType           TenantType `json:"tenantType"`
}
