package anagram

import (
	"reflect"
	"strings"
)

// Detect returns all anagrams of s
func Detect(s string, c []string) []string {

	var anagrams []string
	//counters for characters
	c1 := count(s)
	for _, st := range c {
		// not anagram if different length or the same string
		if len(s) == len(st) && strings.ToLower(s) != strings.ToLower(st) {
			c2 := count(st)
			eq := reflect.DeepEqual(c1, c2)
			if eq {
				anagrams = append(anagrams, st)
			}
		}

	}
	return anagrams
}

func count(s string) map[rune]int {
	c := make(map[rune]int)
	for _, r := range strings.ToLower(s) {
		c[r]++
	}
	return c
}
