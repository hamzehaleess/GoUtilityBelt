package strings

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"racecar", "racecar"},
	}

	for _, test := range tests {
		result := Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"hello", false},
		{"A man a plan a canal Panama", true},
		{"", true},
		{"a", true},
	}

	for _, test := range tests {
		result := IsPalindrome(test.input)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "Hello World"},
		{"HELLO WORLD", "Hello World"},
		{"", ""},
		{"a", "A"},
	}

	for _, test := range tests {
		result := Capitalize(test.input)
		if result != test.expected {
			t.Errorf("Capitalize(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "helloWorld"},
		{"Hello World", "helloWorld"},
		{"HELLO WORLD", "helloWorld"},
		{"", ""},
		{"hello", "hello"},
	}

	for _, test := range tests {
		result := CamelCase(test.input)
		if result != test.expected {
			t.Errorf("CamelCase(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "hello_world"},
		{"HelloWorld", "hello_world"},
		{"helloWorld", "hello_world"},
		{"", ""},
		{"hello", "hello"},
	}

	for _, test := range tests {
		result := SnakeCase(test.input)
		if result != test.expected {
			t.Errorf("SnakeCase(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		suffix   string
		expected string
	}{
		{"hello world", 5, "...", "he..."},
		{"hello", 10, "...", "hello"},
		{"hello world", 8, "", "hello wo"},
	}

	for _, test := range tests {
		result := Truncate(test.input, test.length, test.suffix)
		if result != test.expected {
			t.Errorf("Truncate(%q, %d, %q) = %q; expected %q", test.input, test.length, test.suffix, result, test.expected)
		}
	}
}

func TestRemoveSpaces(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "helloworld"},
		{"  hello  world  ", "helloworld"},
		{"", ""},
		{"nospaces", "nospaces"},
	}

	for _, test := range tests {
		result := RemoveSpaces(test.input)
		if result != test.expected {
			t.Errorf("RemoveSpaces(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello world", 2},
		{"", 0},
		{"hello", 1},
		{"  hello  world  test  ", 3},
	}

	for _, test := range tests {
		result := CountWords(test.input)
		if result != test.expected {
			t.Errorf("CountWords(%q) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"   ", true},
		{"hello", false},
		{" hello ", false},
	}

	for _, test := range tests {
		result := IsEmpty(test.input)
		if result != test.expected {
			t.Errorf("IsEmpty(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestContainsIgnoreCase(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected bool
	}{
		{"Hello World", "hello", true},
		{"Hello World", "WORLD", true},
		{"Hello World", "xyz", false},
		{"", "hello", false},
	}

	for _, test := range tests {
		result := ContainsIgnoreCase(test.input, test.substr)
		if result != test.expected {
			t.Errorf("ContainsIgnoreCase(%q, %q) = %v; expected %v", test.input, test.substr, result, test.expected)
		}
	}
}
