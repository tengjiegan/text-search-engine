# 🔍 Text Search Engine

## Description

A Go library providing efficient custom string processing utilities and substring search capabilities. 
Designed for high-performance text processing, searching and reusability in text-based applications.

### Features
- Optimised text search: time and space efficient algorithm
- Human-friendly indexing: 1-based indexed starting from 1
- Find all matches: Locates all occurrence of search term
- Flexible search options:
    - Exact or partial matches
    - Case sensitive or insensitive
    - Support symbols, punctuation
    - Works with ASCII and Unicode text

### Tech Stacks
- Language: Go v1.24.5
- Testing Framework: Go's built-in `testing` package
- Reflection: `reflect` package for advanced type operations

### Package Reference

#### Package: `pkg/substring`

**`Search(textToSearch, subtext string) string`**

Parameters:
- `textToSearch` (string): The source text to search within
- `subtext` (string): The pattern/substring to find

Returns:
- `string`: Comma-separated string of 1-based positions where matches were found (e.g., "1, 5, 12"), "\<No Output\>" if no matches

#### Package: `internal/utils`

**Utility Functions:**
- `LengthString(input string) int`         - Calculate length of string
- `Lower(s string) string`                 - Convert string to lowercase
- `MaxOf[T cmp.Ordered](a, b T) T`         - Get maximum of two comparable values
- `SliceToString[T any](input []T) string` - Convert slice to comma-separated string
- `LengthOf[T any](input []T) int`         - Calculate length of any slice

#### Package: `internal/constants`

**Constants:**
- `NoOutput`         - Output string when there are no occurrences found
- `TextToSearch`     - Given hard-coded test string 

### Usage Examples

#### Single Text Search
```go
package main

import (
    "fmt"
    "github.com/tengjiegan/textsearchengine/pkg/substring"
)

func main() {
    text := "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
    
    // Case-insensitive search
    result := substring.Search(text, "peter")
    fmt.Printf("'peter' found at positions: %s\n", result)
    // Output: 'peter' found at positions: 1, 20, 75
}
```

#### Multi Text Search
```go
package main

import (
    "fmt"
    "github.com/tengjiegan/textsearchengine/pkg/substring"
)

func main() {
    text := "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
    
    // Multi search with a collection of strings
    searchTerms := []string{"peter", "pick", "pitted"}
    
    for _, term := range searchTerms {
        result := substring.Search(text, term)
        fmt.Printf("'%s' found at positions: %s\n", term, result)
    }
    
    // Output:
    // 'peter' found at positions: 1, 20, 75
    // 'pick' found at positions: 30, 58
    // 'pitted' found at positions: 51
}
```

### Module/Project Structure
```
textsearchengine/
├── go.mod                      # Main module definition
├── README.md                   # Project documentation
├── pkg/                        # Public packages
│   └── substring/              # Core search functionality
│       ├── search.go           # Search functionality
│       └── search_test.go      # Unit tests
└── internal/                   # Internal packages
    └── constants/              # Common constant values
    │   ├── go.mod              # Constants module definition
    │   └── values.go           # Hard coded string values
    └── utils/                  # Common utility functions
        ├── go.mod              # Utils module definition
        ├── common.go           # Generic and helper functions
        └── common_test.go      # Utility function tests
```

### Development

#### Prerequisites
- Go 1.18 or later
- Git for version control

#### Development Tools
- `golangci-lint run` package for linting
- `testing` package for unit tests

### Building and Testing

#### Run Tests
```bash
# Test all packages
go test ./...

# Test with verbose output
go test -v ./...

# Test specific package
go test ./pkg/substring
go test ./internal/utils

# Run with coverage
go test -cover ./...
```

#### Build
```bash
# Build the module
go build ./...

# Verify module
go mod verify

# Tidy dependencies
go mod tidy
```
