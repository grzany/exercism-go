//Package reverse implements simple method to reverse a string
package reverse

//Reverse returns reversed string from given input
func Reverse(s string) string {
	var ns string
	for _, c := range s {
		ns = string(c) + ns
	}
	return ns
}
