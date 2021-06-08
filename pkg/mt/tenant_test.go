package mt_test

import (
	"encoding/json"
	"fmt"
	"github.com/fossteams/teams-api/pkg/models"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetTenants(t *testing.T) {
	userSvc := initTest(t)
	userSvc.DebugSave(true)
	userSvc.DebugDisallowUnknownFields(true)

	tenants, err := userSvc.GetTenants()
	assert.Nil(t, err)
	assert.NotNil(t, tenants)
	assert.GreaterOrEqual(t, 1, len(tenants))
}

func TestParseTenantsResponse(t *testing.T) {
	f, err := os.Open("../../resources/mt/tenants/tenants-1.json")
	defer f.Close()
	if err != nil {
		t.Fatalf("unable to open file: %v", err)
	}

	var tenants []models.Tenant
	dec := json.NewDecoder(f)

	err = dec.Decode(&tenants)
	dec.DisallowUnknownFields()
	if err != nil {
		t.Fatalf("unable to decode JSON: %v", err)
	}
	fmt.Printf("tenants:\n%+v\n", tenants)
	assert.NotNil(t, tenants)
	assert.Equal(t, 1, len(tenants))
	assert.Equal(t, "FossTeams", tenants[0].TenantName)
	assert.Equal(t, "c9fa8756-bafa-47d4-9d21-71f2b67c5e1f", tenants[0].TenantID)
	assert.Equal(t, models.Organization, tenants[0].TenantType)
}

func TestGetVerifiedDomains(t *testing.T) {
	userSvc := initTest(t)

	verifiedDomains, err := userSvc.GetVerifiedDomains()
	assert.Nil(t, err)
	assert.NotNil(t, verifiedDomains)
	assert.GreaterOrEqual(t, len(*verifiedDomains), 0)
}
