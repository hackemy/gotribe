package models

type Image struct {
	Variants map[string]ImageItem `json:"variants"`
}

type ImageItem struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Url    string `json:"url"`
	Name   string `json:"name"`
}
