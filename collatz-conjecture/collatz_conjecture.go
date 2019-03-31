//Package collatzconjecture defines functions for 3x+1 problem calculations
package collatzconjecture

import "errors"

//CollatzConjecture calculates number of steps to reach 1 from any positive integer
func CollatzConjecture(n int) (int, error) {
	var i int
	if n <= 0 {
		return 0, errors.New("Passed a not positive number")
	}
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		i++
	}
	return i, nil
}
