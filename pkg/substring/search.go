package substring

import (
	"github/tengjiegan/internal/constants"
	"github/tengjiegan/internal/utils"
)

// Main search function
// Search all the occurrences of subtext in textToSearch
// Returns all 1-indexed based position of the subtext occurrences
func Search(textToSearch, subtext string) string {
	// Get the length of both strings
	textLen, subtextLen := utils.LengthString(textToSearch), utils.LengthString(subtext)

	// Guard: return early
	if subtextLen == 0 || textLen == 0 || subtextLen > textLen {
		return constants.NoOutput
	}

	// Normalise strings for case-insensitive search
	textToSearch, subtext = utils.Lower(textToSearch), utils.Lower(subtext)

	// Build up the map
	charMap := buildCharMap(subtext)

	// Initialise slice with estimated capacity
	matches := make([]int, 0, textLen/subtextLen+1)

	// Convert strings to rune slices and update length
	textRunes, subtextRunes := []rune(textToSearch), []rune(subtext)
	textLen, subtextLen = utils.LengthOf(textRunes), utils.LengthOf(subtextRunes)

	// Begin search
	for pos := 0; pos <= textLen-subtextLen; {
		mismatchPos := findPattern(textRunes, subtextRunes, pos)
		// Occurrence found, append index and continue search
		if mismatchPos < 0 {
			matches = append(matches, pos+1) // Ensure its 1-indexed
			pos++
		} else {
			// Mismatch found, shift pointer
			pos += calculateShift(
				charMap,
				textRunes[pos+mismatchPos],
				mismatchPos,
				subtextLen,
			)
		}
	}

	// Return results if any matches found
	if utils.LengthOf(matches) > 0 {
		return utils.SliceToString(matches)
	}

	return constants.NoOutput
}

// Creates map for quick lookup of indexes/position
func buildCharMap(subtext string) map[rune]int {
	charMap := make(map[rune]int)
	for i, c := range subtext {
		charMap[c] = i
	}
	return charMap
}

// Returns the mismatch position or -1 if pattern matches
func findPattern(textRunes, subtextRunes []rune, startPos int) int {
	for i := utils.LengthOf(subtextRunes) - 1; i >= 0; i-- {
		if subtextRunes[i] != textRunes[startPos+i] {
			return i
		}
	}
	return -1
}

// Computes the number of positions to shift
func calculateShift(charMap map[rune]int, mismatchChar rune, mismatchPos, subtextLen int) int {
	if lastOccurrence, exists := charMap[mismatchChar]; exists {
		return utils.MaxOf(mismatchPos-lastOccurrence, 1)
	}
	return subtextLen
}
