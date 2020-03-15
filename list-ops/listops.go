// Package listops implements common operations for list of ints
package listops

type binFunc = func(x, y int) int

//IntList defines type of list of ints
type IntList []int

// Foldr inplements right fold
func (list IntList) Foldr(fun binFunc, i int) int {

	return 0
}

// Foldl inplements right fold
func (list IntList) Foldl(fun binFunc, i int) int {

	return 0
}

// Length returns length of the list
func (list IntList) Length() int {
	c := 0
	for _, i := range list {
		c++
	}
	return c
}
