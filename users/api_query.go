package users

import (
	"net/url"

	"github.com/hackemy/gotribe/common"
	"github.com/hackemy/gotribe/models"
)

func Query(accessToken string, query url.Values) ([]models.User, []interface{}, *models.Meta, error) {

	bodyBytes, err := common.GetData(accessToken, "/v1/integration_api/users/query", query)
	if err != nil {
		return nil, nil, nil, err
	}

	items, included, meta, err := models.UnmarshalResponses(bodyBytes)
	if err != nil {
		return nil, nil, nil, err
	}

	results := []models.User{}
	for _, item := range items {
		if tx, ok := item.(models.User); ok {
			results = append(results, tx)
		}
	}

	return results, included, meta, nil
}
