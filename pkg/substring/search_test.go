package substring

import (
	"github/tengjiegan/internal/constants"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected string
	}{
		// Basic functionality
		{"hello world", "world", "7"},
		{"hello world hello", "hello", "1, 13"},
		{"abcdef", "c", "3"},
		{"ababab", "a", "1, 3, 5"},
		{constants.TextToSearch, "Peter", "1, 20, 75"},

		// Case insensitive
		{"Hello World", "hello", "1"},
		{"hello world", "WORLD", "7"},
		{"HeLLo WoRLd", "hElLo", "1"},

		// No matches
		{"hello world", "xyz", constants.NoOutput},
		{"", "hello", constants.NoOutput},
		{"", "", constants.NoOutput},
		{"hi", "hello", constants.NoOutput},
		{constants.TextToSearch, "z", constants.NoOutput},
		{constants.TextToSearch, "Peterz", constants.NoOutput},

		// Exact matches
		{"hello", "hello", "1"},
		{"HELLO", "hello", "1"},
		{constants.TextToSearch, "peter", "1, 20, 75"},

		// Overlapping patterns
		{"aaaa", "aa", "1, 2, 3"},
		{"abababab", "abab", "1, 3, 5"},
		{constants.TextToSearch, "pick", "30, 58"},
		{constants.TextToSearch, "pi", "30, 37, 43, 51, 58"},

		// Special characters
		{"hello world test", " world ", "6"},
		{"hello, world!", ", world!", "6"},
		{"test123test456", "123", "5"},

		// Unicode characters
		{"café résumé naïve", "café", "1"},
		{"café résumé naïve", "résumé", "6"},
		{"café résumé naïve", "naïve", "13"},
	}

	for _, tt := range tests {
		if got := Search(tt.text, tt.pattern); !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("Search(%q, %q) = %v, want %v", tt.text, tt.pattern, got, tt.expected)
		}
	}
}

func TestMultiSubtextSearch(t *testing.T) {
	text := constants.TextToSearch

	// Multi search with a collection of search terms
	tests := []struct {
		name        string
		searchTerms []string
		expected    []string
	}{
		{
			name:        "peter_piper_pitted",
			searchTerms: []string{"peter", "piper", "pitted"},
			expected:    []string{"1, 20, 75", "37", "51"},
		},
		{
			name:        "pick_pi_out",
			searchTerms: []string{"pick", "pi", "out"},
			expected:    []string{"30, 58", "30, 37, 43, 51, 58", "83"},
		},
		{
			name:        "the_he_before",
			searchTerms: []string{"the", "he", "before"},
			expected:    []string{"26", "27, 72, 89", "65"},
		},
		{
			name:        "not_found_terms",
			searchTerms: []string{"xyz", "Peterz", "z"},
			expected:    []string{constants.NoOutput, constants.NoOutput, constants.NoOutput},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, term := range tt.searchTerms {
				result := Search(text, term)
				expected := tt.expected[i]
				if result != expected {
					t.Errorf("Search(%q, %q) = %q, want %q", text, term, result, expected)
				}
			}
		})
	}
}

func TestBuildCharMap(t *testing.T) {
	charMap := buildCharMap("hello")

	tests := []struct {
		char     rune
		expected int
		exists   bool
	}{
		{'h', 0, true},
		{'e', 1, true},
		{'l', 3, true},
		{'o', 4, true},
		{'x', 0, false},
		{'z', 0, false},
	}

	for _, tt := range tests {
		if got, exists := charMap[tt.char]; exists != tt.exists || (exists && got != tt.expected) {
			if tt.exists {
				t.Errorf("charMap[%q] = %d, want %d", tt.char, got, tt.expected)
			} else {
				t.Errorf("charMap[%q] should not exist, but got %d", tt.char, got)
			}
		}
	}
}
