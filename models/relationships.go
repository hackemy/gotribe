package models

type Relationship struct {
	Data RelationshipData `json:"data"`
}

// Struct representing the Data of User Relationship of a Listing
type RelationshipData struct {
	ID   UUID   `json:"id"`
	Type string `json:"type"`
}
