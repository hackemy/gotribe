package models

type GeoCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeoBounds struct {
	NE GeoCoordinates `json:"ne"`
	SW GeoCoordinates `json:"sw"`
}

type GeoCountry struct {
	Name      string `json:"name"`
	ShortCode string `json:"shortCode"`
}

type GeoState struct {
	Name      string `json:"name"`
	ShortCode string `json:"shortCode"`
}

type GeoLocation struct {
	Bounds    GeoBounds  `json:"bounds"`
	City      string     `json:"city"`
	Country   GeoCountry `json:"country"`
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	State     GeoState   `json:"state"`
}
