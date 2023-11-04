package users

import (
	"net/url"
	"strconv"

	"github.com/ruralcoder/gotribe/models"
)

func GetAll(accessToken string, query url.Values) ([]models.User, []interface{}, error) {
	var allUsers []models.User
	var allIncluded []interface{}
	hasNextPage := true
	currentPage := 1
	for hasNextPage {
		query.Set("page", strconv.Itoa(currentPage))

		items, included, meta, err := Query(accessToken, query)
		if err != nil {
			return nil, nil, err
		}

		allUsers = append(allUsers, items...)
		allIncluded = append(allIncluded, included...)

		if currentPage >= meta.TotalPages {
			hasNextPage = false
		} else {
			currentPage++
		}

	}

	return allUsers, allIncluded, nil
}
