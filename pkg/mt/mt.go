package mt

import (
	api "github.com/fossteams/teams-api/pkg"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type MTService struct {
	middleTierUrl *url.URL
	region        api.Region
	token         *api.TeamsToken
	client        *http.Client
}

const MiddleTier = "https://teams.microsoft.com/api/mt/"

func NewMiddleTierService(region api.Region, token *api.TeamsToken) (*MTService, error) {
	svcUrl, err := url.Parse(MiddleTier)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient

	return &MTService{
		middleTierUrl: svcUrl,
		token:         token,
		region:        region,
		client:        client,
	}, nil
}

func (m *MTService)  AuthenticatedRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", api.AuthString(m.token))
	return req, nil
}

func (m *MTService) getEndpoint(path string) *url.URL {
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	endpointUrl, err := m.middleTierUrl.Parse(string(m.region) + "/beta/" + path)
	if err != nil {
		return m.middleTierUrl
	}

	return endpointUrl
}