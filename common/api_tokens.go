package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ruralcoder/gotribe/models"
)

const SHARETRIBE_INTEGRATION = "integ"

// Define the GetAccessToken function
func GetAccessToken(config *models.Config) (string, error) {

	if accessToken != "" && isTokenValid() {
		return accessToken, nil
	}

	apiUrl := fmt.Sprintf("%s/v1/auth/token", ApiBaseUrl)
	data := url.Values{}
	data.Set("client_id", config.ClientId)
	data.Set("grant_type", "client_credentials")
	data.Set("client_secret", config.ClientSecret)
	data.Set("scope", SHARETRIBE_INTEGRATION)

	body := strings.NewReader(data.Encode())

	req, err := http.NewRequest("POST", apiUrl, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	bodyStr, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var accessTokenResponse models.TokenResponse
	err = json.Unmarshal(bodyStr, &accessTokenResponse)
	if err != nil {
		return "", err
	}

	accessToken = accessTokenResponse.AccessToken
	setExpiration(accessTokenResponse.ExpiresIn)

	return accessToken, nil
}

var accessToken string
var accessTokenExpires *time.Time

func InvalidateToken() {
	accessTokenExpires = nil
	accessToken = ""
}

func isTokenValid() bool {
	if accessTokenExpires == nil {
		return false
	}
	return time.Now().Before(*accessTokenExpires)
}

func setExpiration(expiresIn int) {
	now := time.Now()
	t := now.Add(time.Duration(expiresIn) * time.Second)
	accessTokenExpires = &t
}
