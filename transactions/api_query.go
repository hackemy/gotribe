package transactions

import (
	"net/url"

	"github.com/ruralcoder/gotribe/common"
	"github.com/ruralcoder/gotribe/models"
)

func Query(accessToken string, query url.Values) ([]*models.Transaction, []interface{}, *models.Meta, error) {
	bodyBytes, err := common.GetData(accessToken, "/v1/integration_api/transactions/query", query)
	if err != nil {
		return nil, nil, nil, err
	}

	items, included, meta, err := models.UnmarshalResponses(bodyBytes)
	if err != nil {
		return nil, nil, nil, err
	}

	var results []*models.Transaction
	for _, item := range items {
		// Type assertion with comma-ok idiom
		if tx, ok := item.(*models.Transaction); ok {
			results = append(results, tx)
		}
	}

	return results, included, meta, nil
}
