// Package validation provides utility functions for data validation
package validation

import (
	"net"
	"regexp"
	"strconv"
	"strings"
)

// IsEmail validates if a string is a valid email address
func IsEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// IsURL validates if a string is a valid URL
func IsURL(url string) bool {
	urlRegex := regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)
	return urlRegex.MatchString(url)
}

// IsPhone validates if a string is a valid phone number (US format)
func IsPhone(phone string) bool {
	phoneRegex := regexp.MustCompile(`^\+?1?[-.\s]?\(?[0-9]{3}\)?[-.\s]?[0-9]{3}[-.\s]?[0-9]{4}$`)
	return phoneRegex.MatchString(phone)
}

// IsZipCode validates if a string is a valid US ZIP code
func IsZipCode(zip string) bool {
	zipRegex := regexp.MustCompile(`^\d{5}(-\d{4})?$`)
	return zipRegex.MatchString(zip)
}

// IsCreditCard validates if a string is a valid credit card number using Luhn algorithm
func IsCreditCard(number string) bool {
	// Remove spaces and dashes
	number = strings.ReplaceAll(strings.ReplaceAll(number, " ", ""), "-", "")
	
	// Check if all characters are digits
	if !regexp.MustCompile(`^\d+$`).MatchString(number) {
		return false
	}
	
	// Check length (13-19 digits for most cards)
	if len(number) < 13 || len(number) > 19 {
		return false
	}
	
	// Luhn algorithm
	sum := 0
	alternate := false
	
	for i := len(number) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(number[i]))
		
		if alternate {
			digit *= 2
			if digit > 9 {
				digit = digit/10 + digit%10
			}
		}
		
		sum += digit
		alternate = !alternate
	}
	
	return sum%10 == 0
}

// IsIP validates if a string is a valid IP address (IPv4 or IPv6)
func IsIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsAlphanumeric checks if a string contains only alphanumeric characters
func IsAlphanumeric(s string) bool {
	alphanumericRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return alphanumericRegex.MatchString(s)
}

// IsNumeric checks if a string contains only numeric characters
func IsNumeric(s string) bool {
	numericRegex := regexp.MustCompile(`^[0-9]+$`)
	return numericRegex.MatchString(s)
}

// IsAlpha checks if a string contains only alphabetic characters
func IsAlpha(s string) bool {
	alphaRegex := regexp.MustCompile(`^[a-zA-Z]+$`)
	return alphaRegex.MatchString(s)
}

// IsStrongPassword validates if a password meets strong criteria
func IsStrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password)
	
	return hasUpper && hasLower && hasDigit && hasSpecial
}
