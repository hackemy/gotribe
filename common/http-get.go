package common

import (
	"fmt"
	"net/http"
	"net/url"
)

func GetData(accessToken, resource string, query url.Values) ([]byte, error) {

	sharetribeResource := fmt.Sprintf("%s%s", sharetribeBaseUrl, resource)
	req, err := http.NewRequest("GET", sharetribeResource, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.URL.RawQuery = query.Encode()

	return executeRequest(req)
}
