package mt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	api "github.com/fossteams/teams-api/pkg"
	"github.com/fossteams/teams-api/pkg/errors"
	"github.com/fossteams/teams-api/pkg/models"
	"github.com/fossteams/teams-api/pkg/util"
	"io"
	"net/http"
	"net/url"
)

func (m *Service) GetTenants() ([]models.Tenant, error) {
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
		bodyString, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("invalid status code %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("invalid status code %d: resp = %s", resp.StatusCode, string(bodyString))
	}

	jsonReader, err := util.GetJSON(resp, m.debugSave)
	if err != nil {
		return nil, fmt.Errorf("unable to read body: %v", err)
	}

	var tenant []models.Tenant
	decoder := json.NewDecoder(jsonReader)
	if m.debugDisallowUnknownFields {
		decoder.DisallowUnknownFields()
	}
	err = decoder.Decode(&tenant)

	if err != nil {
		return nil, err
	}
	return tenant, nil
}

func (m *Service) GetUser(email string) (*models.User, error) {
	endpointUrl := m.getEndpoint(
		fmt.Sprintf(
			"/users/%s/",
			url.PathEscape(email),
		),
	)

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

	jsonReader, err := util.GetJSON(resp, m.debugSave)
	if err != nil {
		return nil, fmt.Errorf("unable to read body: %v", err)
	}

	var userResp models.UserResponse
	dec := json.NewDecoder(jsonReader)
	if m.debugDisallowUnknownFields {
		dec.DisallowUnknownFields()
	}
	err = dec.Decode(&userResp)

	if err != nil {
		return nil, err
	}
	return &userResp.Value, nil
}

func GetTokenEmail(token *api.TeamsToken) (string, error) {
	if token == nil {
		return "", fmt.Errorf("invalid token provided (nil)")
	}
	claims := token.Inner.Claims
	var email string
	switch claims.(type) {
	case jwt.MapClaims:
		mapClaims := claims.(jwt.MapClaims)
		val, ok := mapClaims["email"]
		if ok {
			email = val.(string)
			return email, nil
		}
		val, ok = mapClaims["upn"]
		if ok {
			email = val.(string)
			return email, nil
		}
		return "", fmt.Errorf("JWT doesn't contain email nor upn")
	default:
		return "", fmt.Errorf("JWT doesn't have MapClaims")
	}
}

func (m *Service) GetMe() (*models.User, error) {
	// Retrieve email from token
	email, err := GetTokenEmail(m.token)
	if err != nil {
		return nil, fmt.Errorf("unable to get email from token: %v", err)
	}
	return m.GetUser(email)
}

func (m *Service) FetchShortProfile(mri ...string) ([]models.User, error) {
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

	jsonReader, err := util.GetJSON(resp, m.debugSave)
	if err != nil {
		return nil, fmt.Errorf("unable to read body: %v", err)
	}

	var userResp models.UsersResponse
	dec := json.NewDecoder(jsonReader)
	if m.debugDisallowUnknownFields {
		dec.DisallowUnknownFields()
	}
	err = dec.Decode(&userResp)

	if err != nil {
		return nil, err
	}
	return userResp.Value, nil
}

func (m *Service) GetProfilePicture(emailOrId string) ([]byte, error) {
	endpointUrl := m.getEndpoint(
		fmt.Sprintf("/users/%s/profilepicture?displayname=aaa",
			url.PathEscape(emailOrId),
		),
	)
	req, err := m.AuthenticatedRequest("GET", endpointUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		bodyString, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("invalid status code %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("invalid status code %d: resp = %s", resp.StatusCode, string(bodyString))
	}

	pictureBytes, err := io.ReadAll(resp.Body)
	// pictureBytes is a B64 representation of the JPG image
	// let's decode it
	picBytes, err := base64.StdEncoding.DecodeString(string(pictureBytes))
	return picBytes, err
}

// TODO: Test and check why it returns a 401
func (m *Service) GetTeamsProfilePicture(emailOrId string) ([]byte, error) {
	endpointUrl := m.getEndpoint(
		fmt.Sprintf("/teams/%s/profilepicturev2",
			url.PathEscape(emailOrId),
		),
	)
	req, err := m.CookieRequest("GET", endpointUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		bodyString, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("invalid status code %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("invalid status code %d: resp = %s", resp.StatusCode, string(bodyString))
	}

	pictureBytes, err := io.ReadAll(resp.Body)
	return pictureBytes, err
}
