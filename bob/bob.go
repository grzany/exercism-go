// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"
import "fmt"

// Hey should have a comment documenting it.
func Hey(remark string) string {

	var answer string

	if isYellQuestion(remark) {
		fmt.Println("YellQuestion ", remark)
		answer = "Calm down, I know what I'm doing!"
	} else if isYell(remark) {
		fmt.Println("Yell", remark)
		answer = "Whoa, chill out!"
	} else if isQuestion(remark) {
		fmt.Println("Question", remark)
		answer = "Sure."
	} else if isNothing(remark) {
		fmt.Println("Nothing", remark)
		answer = "Fine. Be that way!"
	} else {
		answer = "Whatever."
	}

	return answer
}

func isQuestion(remark string) bool {
	remark = strings.TrimSpace(remark)
	if len(remark) > 1 && remark[len(remark)-1:] == "?" {
		return true
	}
	return false
}

func isYell(remark string) bool {
	const alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var counter int
	for _, char := range remark {
		if isLetter(char) {
			counter++
		}
		if isLetter(char) && !strings.Contains(alpha, string(char)) {
			fmt.Println("no shout", string(char))
			return false
		}
	}
	if counter == 0 {
		return false
	}
	return true

}

func isYellQuestion(remark string) bool {
	return isYell(remark) && isQuestion(remark)
}

func isNothing(remark string) bool {
	return len(strings.TrimSpace(remark)) == 0
}

func isLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}
