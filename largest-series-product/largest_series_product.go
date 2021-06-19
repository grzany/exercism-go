package lsproduct

import (
	"errors"
	"strconv"
)

//LargestSeriesProduct calculates largest product of contiguous substring of digits of length span
func LargestSeriesProduct(digit string, span int) (int64, error) {
	var p = int64(0)
	if span < 0 {
		return p, errors.New("span must be greater than zero")
	}
	if span > len(digit) {
		return p, errors.New("span must be smaller than string length")
	}

	for i := 0; i < len(digit)-span+1; i++ {
		sub := digit[i : i+span]
		pt := int64(1)
		for n := 0; n < span; n++ {
			//not a number
			if sub[n] < '0' || sub[n] > '9' {
				return p, errors.New("digits input must only contain digits")
			}
			m, _ := strconv.Atoi(string(sub[n]))
			pt *= int64(m)
		}
		if pt > p {
			p = pt
		}
	}
	return p, nil
}
