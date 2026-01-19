package utils

import (
	"cmp"
	"fmt"
)

// generic:
// Calculate the length of any slice or array
func LengthOf[T any](input []T) int {
	count := 0
	for range input {
		count++
	}
	return count
}

// generic:
// Convert a slice or array to string with comma separated values
func SliceToString[T any](input []T) string {
	if input == nil {
		return ""
	}

	var result []byte
	for i, item := range input {
		if i > 0 {
			result = append(result, ", "...)
		}
		itemStr := fmt.Sprint(item)
		result = append(result, itemStr...)
	}
	return string(result)
}

// int, float and string specific:
// Get the maximum between two values of any ordered type
func MaxOf[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// string specific:
// Convert a given string to lower case
func Lower(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if r >= 'A' && r <= 'Z' {
			runes[i] = r + 32
		}
	}
	return string(runes)
}

// string specific:
// Calculate the length of a string
func LengthString(input string) int {
	count := 0
	for range input {
		count++
	}
	return count
}
