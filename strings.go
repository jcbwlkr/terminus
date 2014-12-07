package terminus

import "strings"

// MaxLineLength returns the length of the longest line in a given string.
func maxLineLength(lines string) int {
	max := 0
	for _, line := range strings.Split(lines, "\n") {
		length := len([]rune(line))
		if length > max {
			max = length
		}
	}

	return max
}
