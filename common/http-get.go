package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetData(accessToken, resource string, query url.Values) ([]byte, error) {

	sharetribeResource := fmt.Sprintf("%s%s", ApiBaseUrl, resource)
	req, err := http.NewRequest("GET", sharetribeResource, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {

		var errResponse ErrorResponse
		if err := json.Unmarshal(bodyBytes, &errResponse); err != nil {
			return nil, err
		}

		if len(errResponse.Errors) > 0 {
			return nil, errResponse.Errors[0].ToError()
		} else {
			return nil, fmt.Errorf("sharetribe error %v", resp.StatusCode)
		}
	}

	return bodyBytes, nil
}
