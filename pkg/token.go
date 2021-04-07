package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

type TeamsTokenType string

const (
	TokenSkype  TeamsTokenType = "skypetoken"
	TokenBearer TeamsTokenType = "Bearer"
)

type RootSkypeToken = TeamsToken
type SkypeToken = TeamsToken

type TeamsToken struct {
	Inner *jwt.Token
	Type  TeamsTokenType
}

func AuthString(token *TeamsToken) string {
	if token == nil {
		return ""
	}

	switch token.Type {
	case TokenSkype:
		return fmt.Sprintf("skypetoken=%s", token.Inner.Raw)
	case TokenBearer:
		return fmt.Sprintf("Bearer %s", token.Inner.Raw)
	}

	return ""
}

func GetSkypeToken() (*SkypeToken, error) {
	authClient := New(http.DefaultClient)
	rootToken, err := GetRootToken()
	if err != nil {
		return nil, fmt.Errorf("unable to get root token: %v", err)
	}
	skypeToken, err := authClient.Authz(rootToken, Refresh)
	if err != nil {
		return nil, fmt.Errorf("unable to get skypeToken: %v", err)
	}

	skypeToken.Type = TokenSkype

	return skypeToken, nil
}

func GetSkypeSpacesToken() (*SkypeToken, error) {
	return getToken("skype")
}

func GetTeamsToken() (*TeamsToken, error) {
	return getToken("teams")
}

func GetChatSvcAggToken() (*TeamsToken, error) {
	return getToken("chatsvcagg")
}

func getToken(tokenType string) (*TeamsToken, error) {
	tokenStr := os.Getenv("MS_TEAMS_" + strings.ToUpper(tokenType) + "_TOKEN")
	if tokenStr == "" {
		// Load from filesystem
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("cannot retrieve user homedir")
		}

		jwtPath := path.Join(homeDir, ".config/fossteams/token-" + tokenType + ".jwt")
		f, err := os.Open(jwtPath)
		if err != nil {
			return nil, fmt.Errorf("unable to open %s: %v", jwtPath, err)
		}

		tokenBytes, err := io.ReadAll(f)
		if err != nil {
			return nil, fmt.Errorf("unable to read JWT token from file: %v", err)
		}

		tokenStr = string(tokenBytes)
	}

	jwtToken, _ := jwt.Parse(tokenStr, nil)
	return &TeamsToken{
		Inner: jwtToken,
		Type:  TokenBearer,
	}, nil
}

func GetToken() (*TeamsToken, error) {
	skypeRootToken, err := GetRootToken()
	if err != nil {
		return nil, fmt.Errorf("unable to get skype root token: %v", err)
	}

	// Call Authz
	client := New(nil)
	skypeToken, err := client.Authz(skypeRootToken, Refresh)
	if err != nil {
		return nil, fmt.Errorf("unable to get Skype Token: %v", err)
	}

	return skypeToken, nil
}

func GetRootToken() (*RootSkypeToken, error) {
	return getToken("skype")
}
