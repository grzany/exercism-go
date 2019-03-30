/*
Space implements age calculation on different planets.
*/
package space

import "fmt"

type Planet string

func Age(s float64, p Planet) float64 {
	AgeOnPlanet := map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	const earthYear = 31557600
	m, ok := AgeOnPlanet[p]
	if ok {
		return (s / earthYear) / m
	}
	fmt.Println("Non existing planet")
	return 0
}
