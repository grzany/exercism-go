// Package raindrops implements raindrop speak
package raindrops

import "strconv"

// Convert translates numbers into raindrop-speak
func Convert(i int) string {

	var s string
	if i%3 == 0 {
		s = "Pling"
	}
	if i%5 == 0 {
		s += "Plang"
	}
	if i%7 == 0 {
		s += "Plong"
	}

	if len(s) > 0 {
		return s
	}

	return strconv.Itoa(i)

}
