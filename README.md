# GoUtils - Comprehensive Go Utility Library

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/goutils)](https://goreportcard.com/report/github.com/yourusername/goutils)
[![GoDoc](https://godoc.org/github.com/yourusername/goutils?status.svg)](https://godoc.org/github.com/yourusername/goutils)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

GoUtils is a comprehensive Go utility library that provides **87 helper functions** across **8 specialized packages** for common programming tasks. Built with Go generics for type safety and zero dependencies (except UUID generation), this library offers essential utilities for array manipulation, string processing, validation, cryptography, file operations, mathematical calculations, type conversion, and time utilities.

## 🚀 Features

- **87 utility functions** across 8 specialized packages
- **Type-safe** with Go generics support
- **Zero dependencies** (minimal external packages)
- **Comprehensive test coverage** (100% tested)
- **Production-ready** code quality
- **MIT License** for open source use

## 📦 Installation

```bash
go get github.com/yourusername/goutils
```

## 🏗️ Package Structure

### Arrays Package (10 functions)
Common array/slice operations for data manipulation.

```go
import "github.com/yourusername/goutils/arrays"

// Check if slice contains element
found := arrays.Contains([]int{1, 2, 3}, 2) // true

// Remove element from slice
result := arrays.Remove([]string{"a", "b", "c"}, "b") // ["a", "c"]

// Get unique elements
unique := arrays.Unique([]int{1, 2, 2, 3, 3}) // [1, 2, 3]

// Sum numeric slice
sum := arrays.Sum([]int{1, 2, 3, 4}) // 10

// Find max/min values
max := arrays.Max([]int{1, 5, 3, 2}) // 5
min := arrays.Min([]int{1, 5, 3, 2}) // 1

// Reverse slice
reversed := arrays.Reverse([]int{1, 2, 3}) // [3, 2, 1]

// Split into chunks
chunks := arrays.Chunk([]int{1, 2, 3, 4, 5}, 2) // [[1, 2], [3, 4], [5]]

// Flatten nested slices
flat := arrays.Flatten([][]int{{1, 2}, {3, 4}}) // [1, 2, 3, 4]

// Shuffle slice randomly
shuffled := arrays.Shuffle([]int{1, 2, 3, 4, 5})
```

### Strings Package (10 functions)
Text processing and string manipulation utilities.

```go
import "github.com/yourusername/goutils/strings"

// Reverse string
reversed := strings.Reverse("hello") // "olleh"

// Convert to camelCase
camel := strings.CamelCase("hello world") // "helloWorld"

// Convert to snake_case
snake := strings.SnakeCase("Hello World") // "hello_world"

// Check if palindrome
isPalindrome := strings.IsPalindrome("racecar") // true

// Capitalize first letter
capitalized := strings.Capitalize("hello") // "Hello"

// Truncate with ellipsis
truncated := strings.Truncate("hello world", 8) // "hello..."

// Remove all spaces
noSpaces := strings.RemoveSpaces("hello world") // "helloworld"

// Count words
wordCount := strings.CountWords("hello world test") // 3

// Check if empty or whitespace
isEmpty := strings.IsEmpty("   ") // true

// Case-insensitive contains
contains := strings.ContainsIgnoreCase("Hello World", "WORLD") // true
```

### Validation Package (10 functions)
Input validation for common data formats.

```go
import "github.com/yourusername/goutils/validation"

// Email validation
isValid := validation.IsEmail("user@example.com") // true

// URL validation
isValid := validation.IsURL("https://example.com") // true

// Phone number validation (US format)
isValid := validation.IsPhone("(555) 123-4567") // true

// ZIP code validation
isValid := validation.IsZipCode("12345") // true

// Credit card validation (Luhn algorithm)
isValid := validation.IsCreditCard("4532015112830366") // true

// IP address validation
isValid := validation.IsIP("192.168.1.1") // true

// Alphanumeric check
isValid := validation.IsAlphanumeric("abc123") // true

// Numeric check
isValid := validation.IsNumeric("12345") // true

// Alphabetic check
isValid := validation.IsAlpha("hello") // true

// Strong password validation
isValid := validation.IsStrongPassword("MyPass123!") // true
```

### Math Package (12 functions)
Mathematical operations and number utilities.

```go
import "github.com/yourusername/goutils/math"

// Absolute value
abs := math.Abs(-5) // 5

// Clamp value between min/max
clamped := math.Clamp(15, 1, 10) // 10

// Greatest Common Divisor
gcd := math.GCD(48, 18) // 6

// Least Common Multiple
lcm := math.LCM(4, 6) // 12

// Prime number check
isPrime := math.IsPrime(17) // true

// Factorial
factorial := math.Factorial(5) // 120

// Fibonacci sequence
fib := math.Fibonacci(10) // 55

// Random integer in range
random := math.RandomInt(1, 100) // random number 1-100

// Power calculation
power := math.Power(2, 8) // 256

// Even/odd checks
isEven := math.IsEven(4) // true
isOdd := math.IsOdd(5) // true

// Round to decimal places
rounded := math.Round(3.14159, 2) // 3.14
```

### Crypto Package (10 functions)
Cryptographic operations and security utilities.

```go
import "github.com/yourusername/goutils/crypto"

// MD5 hash
hash := crypto.MD5Hash("hello") // "5d41402abc4b2a76b9719d911017c592"

// SHA256 hash
hash := crypto.SHA256Hash("hello") // "2cf24dba..."

// Generate UUID
uuid := crypto.GenerateUUID() // "550e8400-e29b-41d4-a716-446655440000"

// Random string generation
random := crypto.RandomString(10) // random 10-char string

// Base64 encoding/decoding
encoded := crypto.Base64Encode("hello world")
decoded := crypto.Base64Decode(encoded)

// Random bytes
bytes := crypto.RandomBytes(16) // 16 random bytes

// Generate secure token
token := crypto.GenerateToken(32) // 32-char secure token

// Password hashing (bcrypt)
hashed := crypto.HashPassword("mypassword")
isValid := crypto.VerifyPassword("mypassword", hashed) // true
```

### File Package (12 functions)
File system operations and utilities.

```go
import "github.com/yourusername/goutils/file"

// Check if file exists
exists := file.Exists("example.txt") // true/false

// Check if directory
isDir := file.IsDir("mydir") // true/false

// Get file size
size := file.GetSize("example.txt") // size in bytes

// Get file extension
ext := file.GetExtension("document.pdf") // ".pdf"

// Read file lines
lines, err := file.ReadLines("example.txt")

// Write lines to file
err := file.WriteLines("output.txt", []string{"line1", "line2"})

// Copy file
err := file.Copy("source.txt", "destination.txt")

// Move/rename file
err := file.Move("old.txt", "new.txt")

// Delete file
err := file.Delete("example.txt")

// Create directory
err := file.CreateDir("newdir")

// List files in directory
files, err := file.ListFiles("mydir")

// List directories
dirs, err := file.ListDirs("mydir")
```

### Convert Package (14 functions)
Type conversion and data transformation utilities.

```go
import "github.com/yourusername/goutils/convert"

// String to int conversion
num, err := convert.ToInt("123") // 123

// Various to string conversion
str := convert.ToString(123) // "123"

// String to bool conversion
flag, err := convert.ToBool("true") // true

// String to float conversion
val, err := convert.ToFloat("3.14") // 3.14

// JSON parsing
var data map[string]interface{}
err := convert.ParseJSON(`{"key":"value"}`, &data)

// Convert to JSON
json, err := convert.ToJSON(data)

// String slice to int slice
ints, err := convert.ToIntSlice([]string{"1", "2", "3"}) // [1, 2, 3]

// Int slice to string slice
strs := convert.ToStringSlice([]int{1, 2, 3}) // ["1", "2", "3"]

// Hex conversion
bytes, err := convert.FromHex("48656c6c6f") // "Hello" in bytes
hex := convert.ToHex([]byte("Hello")) // "48656c6c6f"

// Binary conversion
bytes, err := convert.FromBinary("0100100001101001")
binary := convert.ToBinary([]byte("Hi"))
```

### Time Package (9 functions)
Date and time utilities for various operations.

```go
import "github.com/yourusername/goutils/time"
import "time"

// Format duration human-readable
formatted := time.FormatDuration(90 * time.Second) // "1m 30s"

// Parse duration string
duration, err := time.ParseDuration("1h 30m")

// Check if weekend
isWeekend := time.IsWeekend(time.Now()) // true/false

// Calculate business days between dates
days := time.BusinessDays(startDate, endDate)

// Time ago formatting
ago := time.TimeAgo(pastTime) // "2 hours ago"

// Start/end of day
start := time.StartOfDay(time.Now())
end := time.EndOfDay(time.Now())

// Add business days
newDate := time.AddBusinessDays(time.Now(), 5)

// Leap year check
isLeap := time.IsLeapYear(2024) // true

// Days in month
days := time.DaysInMonth(2024, 2) // 29
```

## 🧪 Testing

Run all tests:

```bash
go test ./...
```

Run tests for specific package:

```bash
go test ./arrays
go test ./strings
go test ./validation
```

Run tests with verbose output:

```bash
go test -v ./...
```

## 📋 Requirements

- Go 1.18+ (for generics support)
- No external dependencies (except `github.com/google/uuid` for UUID generation)

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Add tests for your changes
4. Ensure all tests pass (`go test ./...`)
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- [Documentation](https://godoc.org/github.com/yourusername/goutils)
- [Go Report Card](https://goreportcard.com/report/github.com/yourusername/goutils)
- [Issues](https://github.com/yourusername/goutils/issues)

## ⭐ Star History

If you find this library useful, please consider giving it a star on GitHub!

---

**Made with ❤️ for the Go community**