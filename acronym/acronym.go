// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var acronim []string

	for _, str := range strings.Fields(strings.Replace(s, "-", " ", -1)) {

		acronim = append(acronim, strings.ToUpper(string(str[0])))
	}
	return strings.Join(acronim, "")
}
