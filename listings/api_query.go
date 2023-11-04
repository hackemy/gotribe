package listings

import (
	"net/url"

	"github.com/ruralcoder/gotribe/common"
	"github.com/ruralcoder/gotribe/models"
)

func Query(accessToken string, query url.Values) ([]*models.Listing, []interface{}, *models.Meta, error) {

	bodyBytes, err := common.GetData(accessToken, "/v1/integration_api/listings/query", query)
	if err != nil {
		return nil, nil, nil, err
	}

	items, included, meta, err := models.UnmarshalResponses(bodyBytes)
	if err != nil {
		return nil, nil, nil, err
	}

	results := []*models.Listing{}
	for _, item := range items {
		if tx, ok := item.(*models.Listing); ok {
			results = append(results, tx)
		}
	}

	return results, included, meta, nil
}
