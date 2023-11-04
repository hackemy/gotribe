package listings

import (
	"net/url"

	"github.com/hackemy/gotribe/common"
	"github.com/hackemy/gotribe/models"
)

func Show(accessToken string, query url.Values) (*models.Listing, error) {

	bodyBytes, err := common.GetData(accessToken, "/v1/integration_api/listings/show", query)
	if err != nil {
		return nil, err
	}

	item, _, err := models.UnmarshalResponse(bodyBytes)
	if err != nil {
		return nil, err
	}

	return item.(*models.Listing), nil
}
