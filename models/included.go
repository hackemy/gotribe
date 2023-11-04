package models

type IncludedItem struct {
	ID            UUID                  `json:"id"`
	Type          string                `json:"type"`
	Attributes    IncludedAttributes    `json:"attributes"`
	Relationships IncludedRelationships `json:"relationships"`
}

type IncludedAttributes struct {
	Variants    map[string]ImageVariant `json:"variants"`
	UserBanned  bool                    `json:"banned"`
	UserProfile IncludedUserProfile     `json:"profile"`
}

type ImageVariant struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
	Name   string `json:"name"`
}

type IncludedUserProfile struct {
	DisplayName     string             `json:"displayName"`
	AbbreviatedName string             `json:"abbreviatedName"`
	Bio             string             `json:"bio"`
	PublicData      IncludedPublicData `json:"publicData"`
}

type IncludedPublicData struct {
	IsNonprofit bool   `json:"isNPO"`
	ListingId   string `json:"listingId"`
	ProfileType string `json:"profileType"`
}

type IncludedRelationships struct {
	ProfileImage Relationship `json:"profileImage"`
}
