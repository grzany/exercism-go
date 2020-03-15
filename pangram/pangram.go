//Package pangram implement function that checks if a string is a pangram, i.e. contains all the letters of the alphabet
package pangram

import (
	"strings"
)

//IsPangram checks if a string contains all the letters from alphabet
func IsPangram(s string) bool {
	const a = "abcdefghijklmnopqrstuvwxyz"
	s = strings.Replace(s, " ", "", -1)
	if len(s) == 0 {
		return false
	}
	count := a
	for _, c := range s {
		count = strings.Replace(count, strings.ToLower(string(c)), "", -1)
	}
	return len(count) == 0
}
