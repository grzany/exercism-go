// Package romannumerals implements conversion from arabic to roman for all 0 < numbers < 3001
package romannumerals

import (
	"errors"
	"fmt"
	"sort"
)

var intRom = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

// ToRomanNumeral converts arabic numbers to roman
func ToRomanNumeral(n int) (string, error) {
	const min = 0
	const max = 3000
	var s string

	//sort the intRom map keys
	var keys []int
	for k := range intRom {
		//		fmt.Println(k)
		keys = append(keys, k)
	}
	// Sort intRom index in reverse order
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	// Error if arabic is out of range
	if n <= min || n > max {
		return s, errors.New("Arabic number out of range")
	}
	fmt.Println("Arabic number: ", n)
	for n > 0 {
		v := 0
		for _, k := range keys {
			if k <= n {
				v = k
				break
			}
			v = 1
		}
		s += intRom[v]

		n -= v
	}
	fmt.Println("Roman number: ", s)
	return s, nil
}
