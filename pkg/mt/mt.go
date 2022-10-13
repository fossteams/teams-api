package mt

import (
	api "github.com/fossteams/teams-api/pkg"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Service struct {
	middleTierUrl              *url.URL
	region                     api.Region
	client                     *http.Client
	debugSave                  bool
	debugDisallowUnknownFields bool

	token      *api.TeamsToken
	teamsToken *api.TeamsToken
}

const MiddleTier = "https://teams.microsoft.com/api/mt/"

func NewMiddleTierService(region api.Region, token *api.TeamsToken, teamsToken *api.TeamsToken) (*Service, error) {
	svcUrl, err := url.Parse(MiddleTier)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient

	return &Service{
		middleTierUrl:              svcUrl,
		token:                      token,
		teamsToken:                 teamsToken,
		region:                     region,
		client:                     client,
		debugSave:                  false,
		debugDisallowUnknownFields: false,
	}, nil
}

func (m *Service) DebugSave(flag bool) {
	m.debugSave = flag
}

func (m *Service) DebugDisallowUnknownFields(flag bool) {
	m.debugDisallowUnknownFields = flag
}

func (m *Service) AuthenticatedRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", api.AuthString(m.token))
	return req, nil
}

func (m *Service) CookieRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: "TSAUTHCOOKIE", Value: m.teamsToken.Inner.Raw})
	return req, nil
}

func (m *Service) getEndpoint(path string) *url.URL {
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	endpointUrl, err := m.middleTierUrl.Parse(string(m.region) + "/beta/" + path)
	if err != nil {
		return m.middleTierUrl
	}

	return endpointUrl
}
