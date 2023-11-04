package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (u *UUID) UnmarshalJSON(data []byte) error {
	// First, try to unmarshal the data as a simple string (for the string case)
	var idString string
	err := json.Unmarshal(data, &idString)
	if err == nil {
		*u = UUID(idString)
		return nil
	}

	// If unmarshaling as a string failed, try to unmarshal as an SDKUUID object (for the object case)
	var sdkUUID SDKUUID
	err = json.Unmarshal(data, &sdkUUID)
	if err != nil {
		return err
	}

	*u = UUID(sdkUUID.UUID)
	return nil
}

func (a *Agreement) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch value := v.(type) {
	case string:
		lowerValue := strings.ToLower(value)
		*a = Agreement(lowerValue == "yes" || lowerValue == "true")
	case bool:
		*a = Agreement(value)
	case []interface{}:
		if len(value) > 0 {
			if str, ok := value[0].(string); ok {
				lowerStr := strings.ToLower(str)
				*a = Agreement(lowerStr == "yes" || lowerStr == "true")
			} else if boolValue, ok := value[0].(bool); ok {
				*a = Agreement(boolValue)
			} else {
				fmt.Println("invalid value type in array", value[0])
				*a = Agreement(false)
			}
		}
	default:
		*a = Agreement(false)
	}

	return nil
}
