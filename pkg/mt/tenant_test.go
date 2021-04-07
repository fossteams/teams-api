package mt_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTenants(t *testing.T){
	userSvc := initTest(t)

	tenants, err := userSvc.GetTenants()
	assert.Nil(t, err)
	assert.NotNil(t, tenants)
	assert.GreaterOrEqual(t, 1, len(*tenants))
}

func TestGetVerifiedDomains(t *testing.T){
	userSvc := initTest(t)

	verifiedDomains, err := userSvc.GetVerifiedDomains()
	assert.Nil(t, err)
	assert.NotNil(t, verifiedDomains)
	assert.GreaterOrEqual(t,  len(*verifiedDomains), 0)
}
