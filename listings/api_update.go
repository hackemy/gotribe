package listings

import (
	"encoding/json"

	"github.com/hackemy/gotribe/common"
	"github.com/hackemy/gotribe/models"
)

func Create(accessToken string, listing *models.Listing) error {
	jsonData, err := json.Marshal(listing)
	if err != nil {
		return err
	}

	resource := "/v1/integration_api/listings/create?expand=true"
	return common.PostData(accessToken, resource, jsonData)
}

func Update(accessToken string, payload map[string]interface{}) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resource := "/v1/integration_api/listings/update?expand=true"
	return common.PostData(accessToken, resource, jsonData)
}
