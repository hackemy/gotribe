package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const SHARETRIBE_INTEGRATION = "integ"

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

// Define the GetAccessToken function
func GetAccessToken(clientId, clientSecret string) (*TokenResponse, error) {

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("grant_type", "client_credentials")
	data.Set("client_secret", clientSecret)
	data.Set("scope", SHARETRIBE_INTEGRATION)

	payloadReader := strings.NewReader(data.Encode())

	apiUrl := fmt.Sprintf("%s/v1/auth/token", sharetribeBaseUrl)
	req, err := http.NewRequest("POST", apiUrl, payloadReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("Accept", "application/json")

	responseBytes, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	accessTokenResponse := &TokenResponse{}
	err = json.Unmarshal(responseBytes, &accessTokenResponse)
	if err != nil {
		return nil, err
	}

	return accessTokenResponse, nil
}
