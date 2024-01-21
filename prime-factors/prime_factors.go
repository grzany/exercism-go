package prime

// Factors calculates all prime factors of a natural number
func Factors(n int64) []int64 {
	var factors = []int64{}
	factorise(n, 2, &factors)
	return factors
}

func factorise(n int64, factor int64, factors *[]int64) {
	if n < 2 {
		return
	}
	if n%factor == 0 {
		*factors = append(*factors, factor)
		factorise(n/factor, factor, factors)
	} else {
		factorise(n, factor+1, factors)
	}
}
