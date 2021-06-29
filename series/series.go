package series

// All returns all substrings of s with length n
func All(n int, s string) []string {
	var res []string
	// make sure we don't try to extract substring longer that the string
	if n <= len(s) {
		for i := 0; i+n <= len(s); i++ {
			sub := s[i : n+i]
			res = append(res, sub)
		}
	}
	return res

}

// UnsafeFirst returns the first substring of s with length n
func UnsafeFirst(n int, s string) string {
	var res string
	if n <= len(s) {
		res = s[0:n]
	}
	return res
}
