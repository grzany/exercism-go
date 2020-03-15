// summultiples calculates sum of all unique multiples of numbers in a given limit
package summultiples


// SumMultiples calculates sum of all unique multiples 
func SumMultiples(limit int, divisors ...int) int{
	seen := make(map[int]bool)
	sum := 0
	for _, num := range divisors {
		// don't divide by zero
		if num == 0 {
			continue
		}
		for n := 1; n < limit; n++ {
			if n%num == 0 && !seen[n]{
				sum += n
				//track already used multiples
				seen[n] = true
			}
		}
	}
	return sum
}