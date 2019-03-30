// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

//Package gigasecond provides methods for time conversion
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds 1 gigasecond to the time passed
func AddGigasecond(t time.Time) time.Time {

	t = t.Add(time.Second * 1000000000)
	return t
}
