// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb provides a way to construct proverbs from given string slice
package proverb

import "fmt"

// Proverb generates a string slice with formatted proverb
func Proverb(rhyme []string) []string {

	var proverb []string

	if len(rhyme) == 1 {
		return []string{fmt.Sprintf("And all for the want of a %s.", rhyme[0])}
	}
	if len(rhyme) > 1 {
		for i := 0; i < len(rhyme)-1; i++ {
			proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
		proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}
	return proverb
}
