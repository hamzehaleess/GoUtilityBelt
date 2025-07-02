// Package math provides utility functions for mathematical operations
package math

import (
	"math"
	"math/rand"
	"time"
)

// Abs returns the absolute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Clamp constrains a value between a minimum and maximum
func Clamp[T ~int | ~float64 | ~int32 | ~int64 | ~float32](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// GCD calculates the Greatest Common Divisor of two integers
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM calculates the Least Common Multiple of two integers
func LCM(a, b int) int {
	return Abs(a*b) / GCD(a, b)
}

// IsPrime checks if a number is prime
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Factorial calculates the factorial of a number
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// Fibonacci calculates the nth Fibonacci number
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// RandomInt generates a random integer between min and max (inclusive)
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// Power calculates x raised to the power of y
func Power(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// IsEven checks if a number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// IsOdd checks if a number is odd
func IsOdd(n int) bool {
	return n%2 != 0
}

// Round rounds a float64 to the nearest integer
func Round(x float64) int {
	return int(math.Round(x))
}
