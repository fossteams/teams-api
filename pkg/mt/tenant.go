package mt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type VerifiedDomain struct {
	Name string
}

func (m *MTService) GetVerifiedDomains() (*[]VerifiedDomain, error) {
	endpointUrl := m.getEndpoint("/tenant/verifiedDomains")
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

	var verifiedDomains []VerifiedDomain
	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	err = dec.Decode(&verifiedDomains)

	if err != nil {
		return nil, err
	}
	return &verifiedDomains, nil
}
