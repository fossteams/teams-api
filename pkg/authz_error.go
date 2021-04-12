package api

import "fmt"

const (
	GuestUserNotRedeemed AuthzErrorCode = "GuestUserNotRedeemed"
)

type AuthzErrorCode string

var _ error = AuthzError{}

type AuthzError struct {
	ErrorCode AuthzErrorCode `json:"errorCode"`
	Message   string         `json:"message"`
}

func (a AuthzError) Error() string {
	return fmt.Sprintf("AuthzError: %s - %s", a.ErrorCode, a.Message)
}