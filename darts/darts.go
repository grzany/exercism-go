// Package darts calculates points earned in a single toss in the dart game
package darts

const (
	inner = 1
	middle = 5
	outer = 10
)

//Score calculates dart points for given cartesian coordinates
func Score(x,y float64) int {
	square := x*x + y*y
	score := 0

	if square <= inner * inner  {
		score = 10
		return score
	} else if square <= middle * middle  {
		score = 5
		return score
	} else if square <= outer * outer  {
		score = 1
		return score
	}
	return score
}
