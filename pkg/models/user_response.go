package models

type UsersResponse struct {
	Value []User `json:"value"`
	Type  string `json:"type"`
}
