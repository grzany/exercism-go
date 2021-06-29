package rotationalcipher

// RotationalCipher implements Cesear cipher
func RotationalCipher(s string, key int) string {
	c := make([]byte, len(s))
	var a int
	for k, v := range s {
		// lower case letters
		if v <= 'z' && v >= 'a' {
			a = 'a'
			// upper case letters
		} else if v <= 'Z' && v >= 'A' {
			a = 'A'
			//punctuations
		} else {
			c[k] = s[k]
			continue
		}
		c[k] = byte(a + ((int(v)-a)+key)%26)

	}
	return string(c)
}
