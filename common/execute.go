package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	sharetribeBaseUrl = "https://flex-api.sharetribe.com"
)

func executeRequest(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

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
