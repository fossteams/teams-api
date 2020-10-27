package mt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/fossteams/teams-api/api"
	"io/ioutil"
	"net/http"
	"net/url"
)

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

func (m *MTService) GetTenants() (*[]Tenant, error) {
	endpointUrl := m.getEndpoint("/users/tenants")
	req, err := m.AuthenticatedRequest("GET", endpointUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		bodyString, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("invalid status code %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("invalid status code %d: resp = %s", resp.StatusCode, string(bodyString))
	}

	var tenant []Tenant
	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&tenant)

	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (m *MTService) GetUser(email string) (*User, error) {
	endpointUrl := m.getEndpoint("/users/" + url.PathEscape(email) + "/")
	req, err := m.AuthenticatedRequest("GET", endpointUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, api.InvalidResponseError(resp)
	}

	var userResp UserResponse
	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	err = dec.Decode(&userResp)

	if err != nil {
		return nil, err
	}
	return &userResp.Value, nil
}

type UsersResponse struct {
	Value []User `json:"value"`
	Type  string `json:"type"`
}

func (m *MTService) FetchShortProfile(mri ...string) (*[]User, error) {
	endpointUrl := m.getEndpoint("/users/fetchShortProfile")
	values := endpointUrl.Query()
	values.Add("isMailAddress", "false")
	values.Add("enableGuest", "true")
	values.Add("includeIBBarredUsers", "false")
	values.Add("skypeTeamsInfo", "true")

	endpointUrl.RawQuery = values.Encode()

	bodyBytes, err := json.Marshal(mri)
	if err != nil {
		return nil, err
	}
	bodyBytesReader := bytes.NewReader(bodyBytes)

	req, err := m.AuthenticatedRequest("POST", endpointUrl.String(), bodyBytesReader)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, api.InvalidResponseError(resp)
	}

	var userResp UsersResponse
	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	err = dec.Decode(&userResp)

	if err != nil {
		return nil, err
	}
	return &userResp.Value, nil
}

func (m *MTService) GetProfilePicture(email string) ([]byte, error) {
	endpointUrl := m.getEndpoint("/users/" + url.PathEscape(email) + "/profilepicture?displayname=aaa")
	req, err := m.AuthenticatedRequest("GET", endpointUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		bodyString, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("invalid status code %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("invalid status code %d: resp = %s", resp.StatusCode, string(bodyString))
	}

	pictureBytes, err := ioutil.ReadAll(resp.Body)
	// pictureBytes is a B64 representation of the JPG image
	// let's decode it
	picBytes, err := base64.StdEncoding.DecodeString(string(pictureBytes))
	return picBytes, err
}
