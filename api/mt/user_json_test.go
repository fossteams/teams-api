package mt

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserJson(t *testing.T) {
	const resp = `{
	  "value": {
		"department": "Sales",
		"mobile": "+41-79-1234 56 78",
		"physicalDeliveryOfficeName": "OfficeName",
		"userLocation": "OfficeName",
		"accountEnabled": true,
		"showInAddressList": true,
		"mail": "user@contoso.com",
		"objectType": "User",
		"telephoneNumber": "+41-58-100 23 45",
		"preferredLanguage": "EN-US",
		"skypeTeamsInfo": {
		  "isSkypeTeamsUser": true
		},
		"featureSettings": {
		  "isPrivateChatEnabled": true,
		  "enableShiftPresence": false,
		  "coExistenceMode": "Islands",
		  "enableScheduleOwnerPermissions": false
		},
		"sipProxyAddress": "user@contoso.com",
		"smtpAddresses": [
		  "USERNAME@contoso.mail.onmicrosoft.com",
		  "USERNAME@contoso.onmicrosoft.com",
		  "user@contoso.com"
		],
		"isSipDisabled": false,
		"isShortProfile": false,
		"phones": [
		  {
			"type": "Mobile",
			"number": "+41-79-1234 56 78"
		  }
		],
		"responseSourceInformation": "AAD",
		"userPrincipalName": "user@contoso.com",
		"givenName": "John",
		"surname": "Doe",
		"jobTitle": "Mr.",
		"email": "user@contoso.com",
		"userType": "Member",
		"displayName": "Doe John, Sales",
		"type": "ADUser",
		"mri": "8:orgid:07d8f79b-8071-4488-bb0e-fb8f340315f8",
		"objectId": "07d8f79b-8071-4488-bb0e-fb8f340315f8"
	  },
	  "type": "Microsoft.SkypeSpaces.MiddleTier.Models.AadMember"
	}`
	
	bytesReader := bytes.NewReader([]byte(resp))
	decoder := json.NewDecoder(bytesReader)
	decoder.DisallowUnknownFields()
	
	var userResp UserResponse
	err := decoder.Decode(&userResp)
	assert.Nil(t, err)

	user := userResp.Value
	assert.NotNil(t, user)
	assert.Equal(t, "John", user.GivenName)
	assert.Equal(t, "Doe", user.Surname)
}
