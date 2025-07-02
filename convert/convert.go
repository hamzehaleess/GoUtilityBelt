// Package convert provides utility functions for type conversion
package convert

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// ToInt converts various types to int
func ToInt(value interface{}) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	case string:
		return strconv.Atoi(v)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to int", value)
	}
}

// ToString converts various types to string
func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		return fmt.Sprintf("%v", value)
	}
}

// ToBool converts various types to bool
func ToBool(value interface{}) (bool, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	case string:
		return strconv.ParseBool(v)
	case int:
		return v != 0, nil
	case int32:
		return v != 0, nil
	case int64:
		return v != 0, nil
	case float32:
		return v != 0, nil
	case float64:
		return v != 0, nil
	default:
		return false, fmt.Errorf("cannot convert %T to bool", value)
	}
}

// ToFloat converts various types to float64
func ToFloat(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case string:
		return strconv.ParseFloat(v, 64)
	case bool:
		if v {
			return 1.0, nil
		}
		return 0.0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to float64", value)
	}
}

// ParseJSON parses a JSON string into a map[string]interface{}
func ParseJSON(jsonStr string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	return result, err
}

// ToJSON converts a value to JSON string
func ToJSON(value interface{}) (string, error) {
	bytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToIntSlice converts a slice of interfaces to a slice of ints
func ToIntSlice(values []interface{}) ([]int, error) {
	result := make([]int, len(values))
	for i, value := range values {
		converted, err := ToInt(value)
		if err != nil {
			return nil, fmt.Errorf("error converting index %d: %w", i, err)
		}
		result[i] = converted
	}
	return result, nil
}

// ToStringSlice converts a slice of interfaces to a slice of strings
func ToStringSlice(values []interface{}) []string {
	result := make([]string, len(values))
	for i, value := range values {
		result[i] = ToString(value)
	}
	return result
}

// FromHex converts a hexadecimal string to integer
func FromHex(hex string) (int64, error) {
	return strconv.ParseInt(hex, 16, 64)
}

// ToHex converts an integer to hexadecimal string
func ToHex(value int64) string {
	return strconv.FormatInt(value, 16)
}

// FromBinary converts a binary string to integer
func FromBinary(binary string) (int64, error) {
	return strconv.ParseInt(binary, 2, 64)
}

// ToBinary converts an integer to binary string
func ToBinary(value int64) string {
	return strconv.FormatInt(value, 2)
}
