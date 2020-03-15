//Package luhn implements validation of a string as per Luhn formula
package luhn

import (
	"strings"
)

// Valid checks if a given string per Luhn formula
func Valid(s string) bool {

	const min = 2
	// we don't care about white spaces
	s = strings.ReplaceAll(s, " ", "")
	sum := 0
	//set alt to true for even number of digits
	alt := len(s)%2 == 0

	if len(s) < min {
		return false
	}
	for i := range s {
		d := int(s[i] - '0')
		// only digits are valid
		if d < 0 || d > 9 {
			return false
		}
		// double up every second from the right
		if alt {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		alt = !alt
		sum += d
	}
	return sum%10 == 0
}
