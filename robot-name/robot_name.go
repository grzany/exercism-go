package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var seen_m = map[string]int{}

func (r *Robot) Name() (string, error) {
	const comb = 26 * 26 * 10 * 10 * 10          //maximum number of combination of strings with two upper letters and 3 digits
	if len(r.name) == 0 || seen_m[r.name] == 1 { // new robot
		for len(seen_m) <= comb {
			name := genName()
			_, is_present := seen_m[name]
			if !is_present {
				r.name = name
				seen_m[r.name] = 0
				return r.name, nil
			} else if seen_m[name] == 1 {
				break
			}
		}
		fmt.Printf("len seen %d", len(seen_m))
		return "", errors.New("there are no more names available")
	} else {
		return r.name, nil
	}
}

func (r *Robot) Reset() {
	seen_m[r.name] = 1
	//r.name = genName()
	//r.Name()
}

func genName() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digits = "0123456789"
	const len_l = 2
	const len_d = 3
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	prefix := make([]byte, len_l)
	for i := range prefix {
		prefix[i] = letters[random.Intn(len(letters))]
	}
	suffix := make([]byte, len_d)
	for i := range suffix {
		suffix[i] = digits[random.Intn(len(digits))]
	}
	return string(prefix) + string(suffix)
}
