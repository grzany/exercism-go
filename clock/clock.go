//Package clock implements clock without a data function with add and substract functions
package clock

import (
	"fmt"
)

const minutesInHour = 60
const minutesInDay = 60 * 24

//Clock defines clock type
type Clock struct {
	m int
}

//New is a constructor for a new clock
func New(h int, m int) Clock {
	// number of minutes from given hours and minutes - roll over
	mins := (h*minutesInHour + m) % minutesInDay
	// roll over for negative minutes
	if mins < 0 {
		mins += minutesInDay
	}

	return Clock{mins}
}

//String is a stringifier for Clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.m/minutesInHour, c.m%minutesInHour)
}

//Add implements addition of minutes to the clock
func (c Clock) Add(m int) Clock {
	return New(0, c.m+m)
}

//Subtract implements substruction of minutes from a Clock
func (c Clock) Subtract(m int) Clock {
	return New(0, c.m-m)
}
