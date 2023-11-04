package models

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type FloatValue float64

func (fv *FloatValue) MarshalJSON() ([]byte, error) {
	f := float64(*fv)

	if math.IsInf(float64(f), 1) {
		return json.Marshal("Infinity")
	} else if math.IsInf(f, -1) {
		return json.Marshal("-Infinity")
	} else if math.IsNaN(f) {
		return json.Marshal("NaN")
	}

	return json.Marshal(f)
}

func (iu *FloatValue) UnmarshalJSON(data []byte) error {
	var raw interface{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch v := raw.(type) {
	case float64:
		*iu = FloatValue(v)
	case string:
		cleaned := strings.TrimSpace(v)
		cleaned = strings.ReplaceAll(cleaned, ",", "")
		cleaned = strings.ReplaceAll(cleaned, "$", "")
		cleaned = strings.ToLower(cleaned)

		if cleaned == "" {
			*iu = 0.0 // Default value for empty string
		} else {
			value, err := strconv.ParseFloat(cleaned, 64)
			if err != nil {
				*iu = -1.0 // Default value for non-numeric strings
			} else {
				*iu = FloatValue(value)
			}
		}
	case nil:
		*iu = 0.0 // Default value for nil
	default:
		return fmt.Errorf("unexpected type: %T", v)
	}

	return nil
}
