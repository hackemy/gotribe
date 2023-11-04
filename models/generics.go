package models

import (
	"encoding/json"
	"fmt"
)

type DataResponse struct {
	Data     DataObject   `json:"data"`
	Included []DataObject `json:"included"`
	Meta     Meta         `json:"meta"`
}

type DataResponses struct {
	Data     []DataObject `json:"data"`
	Included []DataObject `json:"included"`
	Meta     Meta         `json:"meta"`
}

type DataObject struct {
	ID            UUID            `json:"id"`
	Type          string          `json:"type"`
	Attributes    json.RawMessage `json:"attributes"`
	Relationships map[string]Data `json:"relationships"`
}

func UnmarshalResponse(responseBytes json.RawMessage) (interface{}, []interface{}, error) {

	var resp DataResponse
	err := json.Unmarshal(responseBytes, &resp)
	if err != nil {
		return nil, nil, err
	}

	item, err := CastObject(resp.Data)
	if err != nil {
		return nil, nil, err
	}

	var includedAttributes []interface{}
	for _, obj := range resp.Included {
		item, err := CastObject(obj)
		if err == nil {
			includedAttributes = append(includedAttributes, item)
		}
	}

	return &item, includedAttributes, nil
}

func UnmarshalResponses(responseBytes json.RawMessage) ([]interface{}, []interface{}, *Meta, error) {

	var resp DataResponses
	err := json.Unmarshal(responseBytes, &resp)
	if err != nil {
		return nil, nil, nil, err
	}

	var dataAttributes []interface{}
	for _, obj := range resp.Data {
		item, err := CastObject(obj)
		if err == nil {
			dataAttributes = append(dataAttributes, item)
		}
	}

	var includedAttributes []interface{}
	for _, obj := range resp.Included {
		item, err := CastObject(obj)
		if err == nil {
			includedAttributes = append(includedAttributes, item)
		}
	}

	return dataAttributes, includedAttributes, &resp.Meta, nil
}

func CastObject(obj DataObject) (interface{}, error) {
	var err error
	switch obj.Type {
	case "booking":
		var b Booking
		err = json.Unmarshal(obj.Attributes, &b)
		return &b, err
	case "image":
		var img Image
		err = json.Unmarshal(obj.Attributes, &img)
		return &img, err
	case "listing":
		var l Listing
		err = json.Unmarshal(obj.Attributes, &l)
		return &l, err
	case "transaction":
		var t Transaction
		err = json.Unmarshal(obj.Attributes, &t)
		return &t, err
	case "user":
		var u User
		err = json.Unmarshal(obj.Attributes, &u)
		return &u, err
	default:
		return nil, fmt.Errorf("sharetribe object %v unsupported", obj.Type)
	}
}
