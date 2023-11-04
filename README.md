# gotribe
Golang library for Sharetribe's APIs based off sharetribe public documentation.

https://www.sharetribe.com/api-reference/integration.html

Currently supports basic operations primarily to retrieve `users`, `listings` and `transactions`

## Installing

```
go get github.com/hackemy/gotribe
```


## Retrieve Access Token 

Retrieve the access token for your application
```
// Get sharetribe access token needed for all subsequent calls
accessToken, err := gotribe.GetAccessToken(config.ClientId, config.ClientSecret)
if err != nil {
    log.Fatal("Error getting access token: ", err)
}
```

## Retrieve Users
```
query := url.Values{
    "pub_isAirBnbListing": {"true"},
}

hasNextPage := true
currentPage := 1
for hasNextPage {
    query.Set("page", strconv.Itoa(currentPage))
    query.Set("perPage", "10")

    items, _, meta, err := users.Query(accessToken, query)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    if len(items) == 0 {
        fmt.Println("No users found")
        return
    }

    for _, user := range items {
        fmt.Println(user)
    }

    if currentPage >= meta.TotalPages {
        hasNextPage = false
    } else {
        currentPage++
        break
    }
}
```


## Retrieve Listings
```
query := url.Values{
    "pub_isAirBnbListing": {"true"},
}

hasNextPage := true
currentPage := 1
for hasNextPage {
    query.Set("page", strconv.Itoa(currentPage))
    query.Set("perPage", "10")

    items, _, meta, err := listings.Query(accessToken, query)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    if len(items) == 0 {
        fmt.Println("No listings found")
        return
    }

    for _, listing := range items {
        fmt.Println(listing)
    }

    if currentPage >= meta.TotalPages {
        hasNextPage = false
    } else {
        currentPage++
        break
    }
}
```

## Retrieve Transactions
```
currentTime := time.Now().UTC()
sevenDaysAgo := currentTime.AddDate(0, 0, -7)
sevenDaysAgoStr := sevenDaysAgo.Format(time.RFC3339)

query := url.Values{
    "createdAtStart": {sevenDaysAgoStr},
}

hasNextPage := true
currentPage := 1
for hasNextPage {
    query.Set("page", strconv.Itoa(currentPage))
    query.Set("perPage", "10")

    items, included, meta, err := transactions.Query(accessToken, query)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    for _, transaction := range items {
        fmt.Println(transaction)
    }

    if currentPage >= meta.TotalPages {
        hasNextPage = false
    } else {
        currentPage++
        break
    }
}
```




