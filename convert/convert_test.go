package convert

import (
	"reflect"
	"testing"
)

func TestToInt(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int
		hasError bool
	}{
		{42, 42, false},
		{int32(42), 42, false},
		{int64(42), 42, false},
		{float32(42.0), 42, false},
		{float64(42.0), 42, false},
		{"42", 42, false},
		{true, 1, false},
		{false, 0, false},
		{"invalid", 0, true},
		{[]int{1, 2, 3}, 0, true},
	}

	for _, test := range tests {
		result, err := ToInt(test.input)
		
		if test.hasError {
			if err == nil {
				t.Errorf("ToInt(%v) expected error but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ToInt(%v) unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("ToInt(%v) = %d; expected %d", test.input, result, test.expected)
			}
		}
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{"hello", "hello"},
		{42, "42"},
		{int32(42), "42"},
		{int64(42), "42"},
		{float32(3.14), "3.14"},
		{float64(3.14), "3.14"},
		{true, "true"},
		{false, "false"},
		{[]int{1, 2, 3}, "[1 2 3]"},
	}

	for _, test := range tests {
		result := ToString(test.input)
		if result != test.expected {
			t.Errorf("ToString(%v) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestToBool(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
		hasError bool
	}{
		{true, true, false},
		{false, false, false},
		{"true", true, false},
		{"false", false, false},
		{"1", true, false},
		{"0", false, false},
		{1, true, false},
		{0, false, false},
		{42, true, false},
		{int32(0), false, false},
		{int64(1), true, false},
		{float32(0.0), false, false},
		{float64(3.14), true, false},
		{"invalid", false, true},
		{[]int{1, 2, 3}, false, true},
	}

	for _, test := range tests {
		result, err := ToBool(test.input)
		
		if test.hasError {
			if err == nil {
				t.Errorf("ToBool(%v) expected error but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ToBool(%v) unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("ToBool(%v) = %v; expected %v", test.input, result, test.expected)
			}
		}
	}
}

func TestToFloat(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected float64
		hasError bool
	}{
		{float64(3.14), 3.14, false},
		{float32(3.14), float64(float32(3.14)), false},
		{42, 42.0, false},
		{int32(42), 42.0, false},
		{int64(42), 42.0, false},
		{"3.14", 3.14, false},
		{true, 1.0, false},
		{false, 0.0, false},
		{"invalid", 0.0, true},
		{[]int{1, 2, 3}, 0.0, true},
	}

	for _, test := range tests {
		result, err := ToFloat(test.input)
		
		if test.hasError {
			if err == nil {
				t.Errorf("ToFloat(%v) expected error but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ToFloat(%v) unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("ToFloat(%v) = %f; expected %f", test.input, result, test.expected)
			}
		}
	}
}

func TestParseJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]interface{}
		hasError bool
	}{
		{`{"name": "John", "age": 30}`, map[string]interface{}{"name": "John", "age": float64(30)}, false},
		{`{"valid": true}`, map[string]interface{}{"valid": true}, false},
		{`{}`, map[string]interface{}{}, false},
		{`invalid json`, nil, true},
	}

	for _, test := range tests {
		result, err := ParseJSON(test.input)
		
		if test.hasError {
			if err == nil {
				t.Errorf("ParseJSON(%q) expected error but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ParseJSON(%q) unexpected error: %v", test.input, err)
			}
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("ParseJSON(%q) = %v; expected %v", test.input, result, test.expected)
			}
		}
	}
}

func TestToJSON(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
		hasError bool
	}{
		{map[string]interface{}{"name": "John", "age": 30}, `{"age":30,"name":"John"}`, false},
		{[]int{1, 2, 3}, `[1,2,3]`, false},
		{"hello", `"hello"`, false},
		{42, `42`, false},
		{true, `true`, false},
	}

	for _, test := range tests {
		result, err := ToJSON(test.input)
		
		if test.hasError {
			if err == nil {
				t.Errorf("ToJSON(%v) expected error but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ToJSON(%v) unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("ToJSON(%v) = %q; expected %q", test.input, result, test.expected)
			}
		}
	}
}

func TestToIntSlice(t *testing.T) {
	tests := []struct {
		input    []interface{}
		expected []int
		hasError bool
	}{
		{[]interface{}{1, 2, 3}, []int{1, 2, 3}, false},
		{[]interface{}{"1", "2", "3"}, []int{1, 2, 3}, false},
		{[]interface{}{1.0, 2.0, 3.0}, []int{1, 2, 3}, false},
		{[]interface{}{true, false}, []int{1, 0}, false},
		{[]interface{}{1, "invalid", 3}, nil, true},
	}

	for _, test := range tests {
		result, err := ToIntSlice(test.input)
		
		if test.hasError {
			if err == nil {
				t.Errorf("ToIntSlice(%v) expected error but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ToIntSlice(%v) unexpected error: %v", test.input, err)
			}
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("ToIntSlice(%v) = %v; expected %v", test.input, result, test.expected)
			}
		}
	}
}

func TestToStringSlice(t *testing.T) {
	tests := []struct {
		input    []interface{}
		expected []string
	}{
		{[]interface{}{1, 2, 3}, []string{"1", "2", "3"}},
		{[]interface{}{"a", "b", "c"}, []string{"a", "b", "c"}},
		{[]interface{}{true, false}, []string{"true", "false"}},
		{[]interface{}{1.5, 2.7}, []string{"1.5", "2.7"}},
	}

	for _, test := range tests {
		result := ToStringSlice(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ToStringSlice(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestFromHex(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
		hasError bool
	}{
		{"ff", 255, false},
		{"FF", 255, false},
		{"100", 256, false},
		{"0", 0, false},
		{"invalid", 0, true},
	}

	for _, test := range tests {
		result, err := FromHex(test.input)
		
		if test.hasError {
			if err == nil {
				t.Errorf("FromHex(%q) expected error but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("FromHex(%q) unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("FromHex(%q) = %d; expected %d", test.input, result, test.expected)
			}
		}
	}
}

func TestToHex(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{255, "ff"},
		{256, "100"},
		{0, "0"},
		{15, "f"},
	}

	for _, test := range tests {
		result := ToHex(test.input)
		if result != test.expected {
			t.Errorf("ToHex(%d) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestFromBinary(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
		hasError bool
	}{
		{"1010", 10, false},
		{"0", 0, false},
		{"1111", 15, false},
		{"invalid", 0, true},
	}

	for _, test := range tests {
		result, err := FromBinary(test.input)
		
		if test.hasError {
			if err == nil {
				t.Errorf("FromBinary(%q) expected error but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("FromBinary(%q) unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("FromBinary(%q) = %d; expected %d", test.input, result, test.expected)
			}
		}
	}
}

func TestToBinary(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{10, "1010"},
		{0, "0"},
		{15, "1111"},
		{1, "1"},
	}

	for _, test := range tests {
		result := ToBinary(test.input)
		if result != test.expected {
			t.Errorf("ToBinary(%d) = %q; expected %q", test.input, result, test.expected)
		}
	}
}
