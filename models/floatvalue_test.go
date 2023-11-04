package models

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedVal FloatValue
		expectErr   bool
	}{
		{
			name:        "Float input",
			input:       "123.45",
			expectedVal: FloatValue(123.45),
			expectErr:   false,
		},
		{
			name:        "Numeric string input",
			input:       `"123.45"`,
			expectedVal: FloatValue(123.45),
			expectErr:   false,
		},
		{
			name:        "Numeric string with currency symbol and comma",
			input:       `"$1,123.45"`,
			expectedVal: FloatValue(1123.45),
			expectErr:   false,
		},
		{
			name:        "Zero point one",
			input:       `"0.1"`,
			expectedVal: FloatValue(0.1),
			expectErr:   false,
		},
		{
			name:        "Empty string",
			input:       `""`,
			expectedVal: FloatValue(0.0),
			expectErr:   false,
		},
		{
			name:        "Non-numeric string",
			input:       `"hello"`,
			expectedVal: FloatValue(-1.0),
			expectErr:   false,
		},
		{
			name:        "Nil value",
			input:       "null",
			expectedVal: FloatValue(0.0),
			expectErr:   false,
		},
		{
			name:        "Unexpected type",
			input:       `{"key": "value"}`,
			expectedVal: FloatValue(0),
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var fv FloatValue
			err := json.Unmarshal([]byte(tt.input), &fv)
			if (err != nil) != tt.expectErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.expectErr)
			}
			if fv != tt.expectedVal {
				t.Errorf("UnmarshalJSON() = %v, want %v", fv, tt.expectedVal)
			}
		})
	}
}
