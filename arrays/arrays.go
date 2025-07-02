// Package arrays provides utility functions for array and slice manipulation
package arrays

import (
	"math/rand"
	"time"
)

// Contains checks if a slice contains a specific element
func Contains[T comparable](slice []T, element T) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

// Remove removes the first occurrence of an element from a slice
func Remove[T comparable](slice []T, element T) []T {
	for i, item := range slice {
		if item == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// Unique returns a new slice with duplicate elements removed
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0)
	
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

// Sum calculates the sum of numeric slice elements
func Sum[T ~int | ~float64 | ~int32 | ~int64 | ~float32](slice []T) T {
	var sum T
	for _, item := range slice {
		sum += item
	}
	return sum
}

// Max returns the maximum value from a numeric slice
func Max[T ~int | ~float64 | ~int32 | ~int64 | ~float32](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	
	max := slice[0]
	for _, item := range slice[1:] {
		if item > max {
			max = item
		}
	}
	return max
}

// Min returns the minimum value from a numeric slice
func Min[T ~int | ~float64 | ~int32 | ~int64 | ~float32](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	
	min := slice[0]
	for _, item := range slice[1:] {
		if item < min {
			min = item
		}
	}
	return min
}

// Reverse returns a new slice with elements in reverse order
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, item := range slice {
		result[len(slice)-1-i] = item
	}
	return result
}

// Chunk splits a slice into chunks of specified size
func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return nil
	}
	
	var chunks [][]T
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// Flatten flattens a 2D slice into a 1D slice
func Flatten[T any](slice [][]T) []T {
	var result []T
	for _, subSlice := range slice {
		result = append(result, subSlice...)
	}
	return result
}

// Shuffle randomly shuffles the elements of a slice
func Shuffle[T any](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	
	return result
}
