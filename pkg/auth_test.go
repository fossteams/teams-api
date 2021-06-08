package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseAuth(t *testing.T) {
	const response = `{"tokens":{"skypeToken":"eyJhbGciOiJSUzI1NiIsImtpZCI6IjEwMiIsInR5cCI6IkpXVCJ9.aaa.eee","expiresIn":86397},"region":"emea","partition":"emea01","regionGtms":{"ams":"https://eu-api.asm.skype.com","amsV2":"https://eu-prod.asyncgw.teams.microsoft.com","amsS2S":"https://eu-storage.asm.skype.com:444","appsDataLayerService":"https://teams.microsoft.com/datalayer/emea","appsDataLayerServiceS2S":"https://deletion-svc-emea.datalayer.teams.microsoft.com","calling_callControllerServiceUrl":"https://api.cc.skype.com","calling_callStoreUrl":"https://api.flightproxy.teams.microsoft.com/api/v2/ep/api.userstore.skype.com/","calling_conversationServiceUrl":"https://api.flightproxy.teams.microsoft.com/api/v2/epconv","calling_keyDistributionUrl":"https://api.flightproxy.teams.microsoft.com/api/v2/ep/api.cc.skype.com/kd","calling_potentialCallRequestUrl":"https://api.flightproxy.teams.microsoft.com/api/v2/ep/api.cc.skype.com/cc/v1/potentialcall","calling_sharedLineOptionsUrl":"https://api.flightproxy.teams.microsoft.com/api/v2/ep/api.cc.skype.com/cc/v1/sharedLineAppearance","calling_udpTransportUrl":"udp://api.flightproxy.teams.microsoft.com:3478","calling_uploadLogRequestUrl":"https://api.flightproxy.teams.microsoft.com/api/v2/ep/api.cc.skype.com/cc/v1/uploadlog/","callingS2S_Broker":"https://api.broker.skype.com","callingS2S_CallController":"https://api.cc.skype.com","callingS2S_CallStore":"https://api.userstore.skype.com/","callingS2S_ContentSharing":"https://api.css.skype.com/contentshare/","callingS2S_ConversationService":"https://api.conv.skype.com/conv/","callingS2S_EnterpriseProxy":"https://api.flightproxy.teams.microsoft.com","callingS2S_MediaController":"https://api.mc.skype.com/media/v2/conversations","callingS2S_PlatformMediaAgent":"https://pma.plat.skype.com:6448/platform/v1/incomingcall","chatService":"https://emea.ng.msg.teams.microsoft.com","chatServiceS2S":"https://emea.pg.msg.infra.teams.microsoft.com","drad":"https://eu.msdrad.skype.com/","mailhookS2S":"https://mailhook.teams.microsoft.com/emea","middleTier":"https://teams.microsoft.com/api/mt/emea","middleTierS2S":"https://teams.microsoft.com/api/mt/emea","mtImageService":"https://teams.microsoft.com/api/mt/emea","powerPointStateService":"https://emea.pptservicescast.officeapps.live.com","search":"https://eu-prod.asyncgw.teams.microsoft.com/msgsearch","searchTelemetry":"https://eu-prod.asyncgw.teams.microsoft.com/msgsearch","teamsAndChannelsService":"https://teams.microsoft.com/api/mt/emea","teamsAndChannelsProvisioningService":"https://teams.microsoft.com/fabric/emea/templates/api","urlp":"https://urlp.asm.skype.com","urlpV2":"https://eu-prod.asyncgw.teams.microsoft.com/urlp","unifiedPresence":"https://presence.teams.microsoft.com","userEntitlementService":"https://teams.microsoft.com/api/ues/emea","userIntelligenceService":"https://teams.microsoft.com/api/nss/emea","userProfileService":"https://teams.microsoft.com/api/userprofilesvc/emea","userProfileServiceS2S":"https://userprofilesvc-emea.teams.microsoft.com","amdS2S":"https://eu-distr.asm.skype.com:444","chatServiceAggregator":"https://chatsvcagg.teams.microsoft.com"},"regionSettings":{"isUnifiedPresenceEnabled":true,"isOutOfOfficeIntegrationEnabled":true,"isContactMigrationEnabled":true,"isAppsDiscoveryEnabled":true,"isFederationEnabled":true},"licenseDetails":{"isFreemium":false,"isBasicLiveEventsEnabled":true,"isTrial":false,"isAdvComms":false}}`
	strReader := bytes.NewReader([]byte(response))
	dec := json.NewDecoder(strReader)
	dec.DisallowUnknownFields()

	var authResp AuthzResponse
	err := dec.Decode(&authResp)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	fmt.Printf("%v", authResp)
}

func TestRefreshToken(t *testing.T) {
	authzClient := New(nil)
	rootToken, err := GetRootToken()
	if err != nil {
		t.Fatalf("unable to get root token: %v", err)
	}
	skypeJwt, err := authzClient.Authz(rootToken, AuthzRefresh)
	if err != nil {
		t.Fatalf("unable to get refresh token: %v", err)
	}

	fmt.Printf("got token=%+v", skypeJwt.Inner.Claims)
}
