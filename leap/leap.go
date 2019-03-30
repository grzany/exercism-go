// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap determines if the year passes is a leap year
package leap

// IsLeapYear returns true|false depending if the year passed is leap or not
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			return false
		} else {
			return true
		}
		return true
	}
	return false
}
