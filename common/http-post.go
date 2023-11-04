package common

import (
	"bytes"
	"fmt"
	"net/http"
)

func PostData(accessToken, resource string, jsonData []byte) error {

	sharetribeResource := fmt.Sprintf("%s%s", sharetribeBaseUrl, resource)

	payloadReader := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest("POST", sharetribeResource, payloadReader)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	_, err = executeRequest(req)

	return err
}
