package api

import (
	"encoding/json"
	"fmt"
	"github.com/fossteams/teams-api/pkg/errors"
	"github.com/fossteams/teams-api/pkg/models"
	"io"
	"net/http"
	"net/url"
)

const SkypeSpacesEndpoint = "https://api.teams.skype.com"

type SkypeSpaceSvc struct {
	token *SkypeToken
	endpoint *url.URL
	httpClient *http.Client
}

func NewSkypeSpaceService(token *SkypeToken) SkypeSpaceSvc {
	endpoint, err := url.Parse(SkypeSpacesEndpoint)
	if err != nil {
		panic(fmt.Sprintf("unable to parse SkypeSpaceService endpoint: %v", err))
	}
	return SkypeSpaceSvc{
		token:    token,
		endpoint: endpoint,
		httpClient: http.DefaultClient,
	}
}

func (s SkypeSpaceSvc) authenticateRequest(req *http.Request) {
	if req == nil {
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token.Inner.Raw))
}

func (s SkypeSpaceSvc) get(path string) (*http.Response, error) {
	if s.httpClient == nil {
		return nil, fmt.Errorf("httpClient is nil")
	}

	theUrl, err := s.endpoint.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("unable to parse URL: %v", err)
	}
	req, err := http.NewRequest(http.MethodGet, theUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	s.authenticateRequest(req)
	return s.httpClient.Do(req)
}

func (s SkypeSpaceSvc) GetTenants() ([]models.Tenant, error) {
	res, err := s.get("/beta/users/tenants")
	if err != nil {
		return nil, err
	}

	expectedStatusCode := http.StatusOK
	if res.StatusCode != expectedStatusCode {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, errors.NewHTTPError(res.StatusCode, expectedStatusCode, nil)
		}
		return nil, errors.NewHTTPError(res.StatusCode, expectedStatusCode, bodyBytes)
	}

	var tenants []models.Tenant
	d := json.NewDecoder(res.Body)
	err = d.Decode(&tenants)

	return tenants, err
}