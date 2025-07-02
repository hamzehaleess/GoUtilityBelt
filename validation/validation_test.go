package validation

import "testing"

func TestIsEmail(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"test@example.com", true},
		{"user.name@domain.co.uk", true},
		{"invalid.email", false},
		{"@example.com", false},
		{"test@", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsEmail(test.input)
		if result != test.expected {
			t.Errorf("IsEmail(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestIsURL(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"https://www.example.com", true},
		{"http://example.com", true},
		{"https://sub.domain.com/path?query=1", true},
		{"ftp://example.com", false},
		{"not-a-url", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsURL(test.input)
		if result != test.expected {
			t.Errorf("IsURL(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestIsPhone(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"(555) 123-4567", true},
		{"555-123-4567", true},
		{"555.123.4567", true},
		{"5551234567", true},
		{"1-555-123-4567", true},
		{"invalid-phone", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsPhone(test.input)
		if result != test.expected {
			t.Errorf("IsPhone(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestIsZipCode(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"12345", true},
		{"12345-6789", true},
		{"1234", false},
		{"123456", false},
		{"abcde", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsZipCode(test.input)
		if result != test.expected {
			t.Errorf("IsZipCode(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestIsCreditCard(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"4532015112830366", true},  // Valid Visa
		{"5555555555554444", true},  // Valid MasterCard
		{"378282246310005", true},   // Valid American Express
		{"1234567890123456", false}, // Invalid
		{"4532-0151-1283-0366", true}, // Valid with dashes
		{"", false},
	}

	for _, test := range tests {
		result := IsCreditCard(test.input)
		if result != test.expected {
			t.Errorf("IsCreditCard(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestIsIP(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"192.168.1.1", true},
		{"255.255.255.255", true},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"::1", true},
		{"256.1.1.1", false},
		{"not-an-ip", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsIP(test.input)
		if result != test.expected {
			t.Errorf("IsIP(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestIsAlphanumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc123", true},
		{"ABC123", true},
		{"abc", true},
		{"123", true},
		{"abc-123", false},
		{"abc 123", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsAlphanumeric(test.input)
		if result != test.expected {
			t.Errorf("IsAlphanumeric(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"0", true},
		{"123.45", false},
		{"abc", false},
		{"12a", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsNumeric(test.input)
		if result != test.expected {
			t.Errorf("IsNumeric(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestIsStrongPassword(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Password123!", true},
		{"StrongP@ss1", true},
		{"password", false},      // No uppercase, digits, or special chars
		{"PASSWORD", false},      // No lowercase, digits, or special chars
		{"Pass123", false},       // No special chars
		{"Password!", false},     // No digits
		{"Short1!", false},       // Too short
		{"", false},
	}

	for _, test := range tests {
		result := IsStrongPassword(test.input)
		if result != test.expected {
			t.Errorf("IsStrongPassword(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
