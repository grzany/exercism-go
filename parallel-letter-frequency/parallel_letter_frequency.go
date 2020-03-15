package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency call concurrently Frequency and combines results
func ConcurrentFrequency(sa []string) FreqMap {
	m := FreqMap{}
	ch := make(chan FreqMap, 5)

	for _, s := range sa {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}
	for range sa {
		for k, v := range <-ch {
			m[k] += v
		}
	}
	return m
}
