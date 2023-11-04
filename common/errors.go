package common

import "fmt"

const (
	codeInvalidToken = "auth-invalid-access-token"
)

type SharetribeError struct {
	ID      string `json:"id"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Title   string `json:"title"`
	Details string `json:"details,omitempty"`
	Source  struct {
		Path []string `json:"path"`
		Type string   `json:"type"`
	} `json:"source"`
}

type ErrorResponse struct {
	Errors []*SharetribeError `json:"errors"`
}

func (obj *SharetribeError) ToError() error {

	msg := obj.Title
	if obj.Details != "" {
		msg = fmt.Sprintf("%v - %v", obj.Title, obj.Details)
	}
	return fmt.Errorf("[Sharetribe] %v - %v", obj.Code, msg)
}

func (obj *SharetribeError) IsInvalidToken() bool {
	return obj.Code == codeInvalidToken
}
