//Package accumulate implements operation which given a collection,
//returns a new collection containing the result of applying that operation to each element of the input collection.
package accumulate

type fn func(string) string

//Accumulate calls a converter passed as f and returns slice of converted strings
func Accumulate(s []string, f fn) []string {
	var newStrings []string
	for _, str := range s {
		newStrings = append(newStrings, f(str))
	}
	return newStrings
}
