package listings

import (
	"net/url"
	"strconv"

	"github.com/ruralcoder/gotribe/models"
)

func GetAll(accessToken string, query url.Values) ([]models.Listing, []interface{}, error) {
	var allListings []models.Listing
	var allIncluded []interface{}

	hasNextPage := true
	currentPage := 1
	for hasNextPage {
		query.Set("page", strconv.Itoa(currentPage))

		items, included, meta, err := Query(accessToken, query)
		if err != nil {
			return nil, nil, err
		}

		allListings = append(allListings, items...)
		allIncluded = append(allIncluded, included...)

		if currentPage >= meta.TotalPages {
			hasNextPage = false
		} else {
			currentPage++
		}
	}

	return allListings, allIncluded, nil
}
