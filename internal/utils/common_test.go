package utils

import (
	"testing"
)

func TestLower(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HELLO", "hello"},
		{"HeLLo WoRLd", "hello world"},
		{"A", "a"},
		{"Hello123", "hello123"},
		{"Hello, World!", "hello, world!"},
	}

	for _, tt := range tests {
		if got := Lower(tt.input); got != tt.expected {
			t.Errorf("Lower(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

func TestLengthString(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"a", 1},
		{"abc123", 6},
		{"hello, world!", 13},
		{"   ", 3},
	}

	for _, tt := range tests {
		if got := LengthString(tt.input); got != tt.expected {
			t.Errorf("LengthString(%q) = %d, want %d", tt.input, got, tt.expected)
		}
	}
}

func TestMaxOf(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 3, 5},
		{5, 5, 5},
		{0, -5, 0},
		{-3, -5, -3},
		{-3, 5, 5},
	}

	for _, tt := range tests {
		if got := MaxOf(tt.a, tt.b); got != tt.expected {
			t.Errorf("MaxOf(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
		}
	}
}

func TestLengthOf(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected int
	}{
		{"empty int slice", []int{}, 0},
		{"int slice", []int{1, 2, 3, 4, 5}, 5},
		{"string slice", []string{"hello", "world"}, 2},
		{"single element", []int{42}, 1},
		{"empty string slice", []string{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.input.(type) {
			case []int:
				if got := LengthOf(v); got != tt.expected {
					t.Errorf("LengthOf(%v) = %d, want %d", v, got, tt.expected)
				}
			case []string:
				if got := LengthOf(v); got != tt.expected {
					t.Errorf("LengthOf(%v) = %d, want %d", v, got, tt.expected)
				}
			}
		})
	}
}

func TestSliceToString(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{"empty int slice", []int{}, ""},
		{"single int", []int{1}, "1"},
		{"multiple ints", []int{1, 2, 3}, "1, 2, 3"},
		{"string slice", []string{"hello", "world"}, "hello, world"},
		{"nil slice", ([]int)(nil), ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.input.(type) {
			case []int:
				if got := SliceToString(v); got != tt.expected {
					t.Errorf("SliceToString(%v) = %q, want %q", v, got, tt.expected)
				}
			case []string:
				if got := SliceToString(v); got != tt.expected {
					t.Errorf("SliceToString(%v) = %q, want %q", v, got, tt.expected)
				}
			}
		})
	}
}
