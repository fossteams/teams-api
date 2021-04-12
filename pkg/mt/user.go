package mt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/fossteams/teams-api/pkg/errors"
	"github.com/fossteams/teams-api/pkg/models"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (m *MTService) GetTenants() ([]models.Tenant, error) {
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

	var tenant []models.Tenant
	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&tenant)

	if err != nil {
		return nil, err
	}
	return tenant, nil
}

func (m *MTService) GetUser(email string) (*models.User, error) {
	endpointUrl := m.getEndpoint("/users/" + url.PathEscape(email) + "/")

	values := endpointUrl.Query()
	values.Add("throwIfNotFound", "false")
	values.Add("isMailAddress", "true")
	values.Add("enableGuest", "true")
	values.Add("includeIBBarredUsers", "true")
	values.Add("skypeTeamsInfo", "true")
	endpointUrl.RawQuery = values.Encode()

	req, err := m.AuthenticatedRequest("GET", endpointUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		return nil, errors.NewHTTPError(expectedStatusCode, resp.StatusCode, nil)
	}

	var userResp models.UserResponse
	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	err = dec.Decode(&userResp)

	if err != nil {
		return nil, err
	}
	return &userResp.Value, nil
}

func (m *MTService) GetMe() (*models.User, error) {
	// Retrieve email from token
	claims := m.token.Inner.Claims
	var email string
	switch claims.(type) {
	case jwt.MapClaims:
		email = claims.(jwt.MapClaims)["upn"].(string)
	default:
		return nil, fmt.Errorf("JWT token doesn't have MapClaims")
	}
	return m.GetUser(email)
}

type UsersResponse struct {
	Value []models.User `json:"value"`
	Type  string `json:"type"`
}

func (m *MTService) FetchShortProfile(mri ...string) ([]models.User, error) {
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
	req.Header.Add("Content-Type", "application/json")
	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		return nil, errors.NewHTTPError(expectedStatusCode, resp.StatusCode, nil)
	}

	var userResp UsersResponse
	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	err = dec.Decode(&userResp)

	if err != nil {
		return nil, err
	}
	return userResp.Value, nil
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
