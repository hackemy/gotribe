package models

type Listing struct {
	Description      string                 `json:"description"`
	Deleted          bool                   `json:"deleted"`
	Geolocation      GeoCoordinates         `json:"geolocation"`
	CreatedAt        string                 `json:"createdAt"`
	State            string                 `json:"state"`
	PrivateData      interface{}            `json:"privateData"`
	Title            string                 `json:"title"`
	AvailabilityPlan map[string]interface{} `json:"availabilityPlan"`
	PublicData       interface{}            `json:"publicData"`
	Price            Price                  `json:"price"`
	Metadata         interface{}            `json:"metadata"`
}
