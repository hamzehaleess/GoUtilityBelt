// Package strings provides utility functions for string manipulation
package strings

import (
	"regexp"
	"strings"
	"unicode"
)

// Reverse returns a string with characters in reverse order
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome checks if a string is a palindrome (ignoring case and spaces)
func IsPalindrome(s string) bool {
	cleaned := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return cleaned == Reverse(cleaned)
}

// Capitalize capitalizes the first letter of each word in a string
func Capitalize(s string) string {
	return strings.Title(strings.ToLower(s))
}

// CamelCase converts a string to camelCase
func CamelCase(s string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return s
	}
	
	result := strings.ToLower(words[0])
	for _, word := range words[1:] {
		if len(word) > 0 {
			result += strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return result
}

// SnakeCase converts a string to snake_case
func SnakeCase(s string) string {
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snake := re.ReplaceAllString(s, `${1}_${2}`)
	return strings.ToLower(strings.ReplaceAll(snake, " ", "_"))
}

// Truncate truncates a string to a specified length with optional suffix
func Truncate(s string, length int, suffix string) string {
	if len(s) <= length {
		return s
	}
	return s[:length-len(suffix)] + suffix
}

// RemoveSpaces removes all whitespace characters from a string
func RemoveSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

// CountWords counts the number of words in a string
func CountWords(s string) int {
	return len(strings.Fields(s))
}

// IsEmpty checks if a string is empty or contains only whitespace
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// Contains checks if a string contains a substring (case-insensitive)
func ContainsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
