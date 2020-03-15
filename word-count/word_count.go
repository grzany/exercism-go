//Package wordcount implements function to calculate number of occurances of a word in a sentence
package wordcount

import (
	"strings"
)

//Frequency defines map to keep word occurences count
type Frequency map[string]int

//WordCount implements calculation of word occurances in a sentense
func WordCount(p string) Frequency {

	var f Frequency
	f = make(map[string]int)
	replacer := strings.NewReplacer(",", " ", ".", "", ";", "", "!", " ", "&", "", "@", "", "$", "", "%", "", "^", "", ":", "")
	p = replacer.Replace(p)
	s := strings.Fields(p)

	for _, w := range s {
		w = strings.ToLower(w)
		w := strings.Trim(w, `'`)
		f[w] = f[w] + 1
	}
	return f
}
