package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func PostData(accessToken, resource string, jsonData []byte) error {

	sharetribeResource := fmt.Sprintf("%s%s", ApiBaseUrl, resource)

	req, err := http.NewRequest("POST", sharetribeResource, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {

		var errResponse ErrorResponse
		if err := json.Unmarshal(bodyBytes, &errResponse); err != nil {
			return err
		}

		return fmt.Errorf(errResponse.Errors[0].Details)
	}

	return nil
}
