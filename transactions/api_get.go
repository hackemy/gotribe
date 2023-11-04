package transactions

import (
	"net/url"
	"strconv"

	"github.com/hackemy/gotribe/models"
)

func GetAll(accessToken string, query url.Values) ([]*models.Transaction, error) {
	var allObjects []*models.Transaction

	hasNextPage := true
	currentPage := 1
	for hasNextPage {
		query.Set("page", strconv.Itoa(currentPage))

		items, _, meta, err := Query(accessToken, query)
		if err != nil {
			return nil, err
		}

		allObjects = append(allObjects, items...)

		if currentPage >= meta.TotalPages {
			hasNextPage = false
		} else {
			currentPage++
		}
	}

	return allObjects, nil
}
