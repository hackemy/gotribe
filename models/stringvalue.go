package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type StringValue string

func (iu *StringValue) UnmarshalJSON(data []byte) error {
	var raw interface{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch v := raw.(type) {
	case int64:
	case float64:
		*iu = StringValue(fmt.Sprintf("%v", v))
	case string:
		*iu = StringValue(strings.TrimSpace(v))
	case nil:
	default:
		fmt.Printf("unexpected type for impactUnits: %v (type: %T)\n", v, v)
		*iu = StringValue("")
	}

	return nil
}
