//Package isogram provides tools to check if a word is an isogram ("nopattern word")
package isogram

import (
	"strings"
)

// IsIsogram checks a strings for being an isogram. White space and hyphen can repeat
func IsIsogram(s string) bool {

	seen := make(map[rune]bool)

	for _, r := range strings.ToLower(s) {
		if r == ' ' || r == '-' {
			continue
		}
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
