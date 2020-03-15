//Package atbash implements atbash cipher
package atbash

import "strings"
import "unicode"

const length = 5

//Atbash implements encryption with atbash cipher
func Atbash(s string) string {
	//tidy up input, no spaces or punctuations
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, ".", "", -1)
	s = strings.ToLower(s)
	var conv, enc string
	var counter = 0
	
	for c := range s {
		counter++
		if !unicode.IsDigit(rune(s[c])) {
			conv = string(rune('z'-s[c]+'a'))
		} else {
			conv = string(rune(s[c]))
		}
		enc = enc + conv
		if counter == length {
			enc = enc + " "
			counter = 0
		}
		//fmt.Println(enc)
	}
	enc = strings.TrimSpace(enc)

	return enc
}
