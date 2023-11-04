package gotribe

import (
	"time"

	"github.com/hackemy/gotribe/common"
)

// Define the GetAccessToken function
func GetAccessToken(clientId, clientSecret string) (string, error) {

	if accessToken != "" && isTokenValid() {
		return accessToken, nil
	}

	tokenResponse, err := common.GetAccessToken(clientId, clientSecret)
	if err != nil {
		return "", err
	}

	accessToken = tokenResponse.AccessToken
	setExpiration(tokenResponse.ExpiresIn)

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
