package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
)

type TeamsTokenType string

const (
	TokenSkype  TeamsTokenType = "skypetoken"
	TokenBearer TeamsTokenType = "Bearer"
)

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

func GetTokenFromEnv() (*TeamsToken, error) {
	tokenStr := os.Getenv("MS_TEAMS_TOKEN")
	skypeTokenStr := os.Getenv("MS_TEAMS_SKYPETOKEN")
	if tokenStr == "" && skypeTokenStr == "" {
		return nil, fmt.Errorf("you must provide a teams token (env: MS_TEAMS_TOKEN) " +
			"or a skype token (MS_TEAMS_SKYPETOKEN)")
	}

	if tokenStr != "" {
		jwtToken, _ := jwt.Parse(tokenStr, nil)
		return &TeamsToken{
			Inner: jwtToken,
			Type:  TokenBearer,
		}, nil
	}

	jwtToken, _ := jwt.Parse(skypeTokenStr, nil)
	// Ignore errors
	return &TeamsToken{
		Inner: jwtToken,
		Type:  TokenSkype,
	}, nil
}
