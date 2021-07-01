// Package listops implements common operations for list of ints
package listops

type binFunc = func(x, y int) int
type predFunc = func(x int) bool
type unaryFunc = func(x int) int

//IntList defines type of list of ints
type IntList []int

// Foldr inplements right fold
func (list IntList) Foldr(fun binFunc, i int) int {

	res := i
	for k := len(list) - 1; k >= 0 && len(list) > 0; k-- {
		if res != 0 {
			res = fun(list[k], res)
		}
	}
	return res
}

// Foldl inplements left fold
func (list IntList) Foldl(fun binFunc, i int) int {
	res := i
	for _, ll := range list {
		if res != 0 {
			res = fun(ll, res)
		}
	}
	return res
}

// Length returns length of the list
func (list IntList) Length() int {
	return len(list)
}

// Filter returns int list filtering on pred function
func (list IntList) Filter(fun predFunc) IntList {
	l := IntList{}
	for _, ll := range list {
		if fun(ll) {
			l = append(l, ll)
		}
	}
	return l
}

// Map returns int list with fun applied on list elements
func (list IntList) Map(fun unaryFunc) IntList {
	l := IntList{}
	for _, ll := range list {
		l = append(l, fun(ll))
	}
	return l
}

// Reverse returns reversed list
func (list IntList) Reverse() IntList {
	ll := IntList{}
	for i := len(list) - 1; i >= 0 && len(list) > 0; i-- {
		ll = append(ll, list[i])
	}
	return ll
}

// Append adds all elements of l to the list
func (list IntList) Append(l IntList) IntList {
	return append(list, l...)

}

// Concat adds to the lists all lists from l
func (list IntList) Concat(l []IntList) IntList {
	for _, ll := range l {
		list = list.Append(ll)
	}
	return list
}
