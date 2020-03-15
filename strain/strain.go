//Package strain implements implements discard and keep functions on collections.
package strain

//Ints defines a slice of ints
type Ints []int

//Lists defines a list of int slices
type Lists [][]int

//Strings defines a slice of strings
type Strings []string

type intfunc func(int) bool
type listfunc func([]int) bool
type stringfunc func(string) bool

//Keep returns a new collections of all those int elements where the predicate is true
func (il Ints) Keep(f intfunc) Ints {
	var in Ints
	for _, i := range il {
		if f(i) {
			in = append(in, i)
		}
	}
	return in
}

//Discard returns a new collections of int elements where the predicate is false
func (il Ints) Discard(f intfunc) Ints {
	var in Ints
	for _, i := range il {
		if !f(i) {
			in = append(in, i)
		}
	}
	return in
}

//Keep returns a new collections of list elements where the predicate is true
func (ll Lists) Keep(f listfunc) Lists {
	li := Lists{}
	for _, l := range ll {
		if f(l) {
			li = append(li, l)
		}
	}
	return li
}

//Keep returns a new collections of string elemets where the predicate is true
func (sl Strings) Keep(f stringfunc) Strings {
	si := Strings{}
	for _, s := range sl {
		if f(s) {
			si = append(si, s)
		}
	}
	return si
}
