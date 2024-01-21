// Package weather provides mechanisms to forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition defines variable to hold current weather condition.
var CurrentCondition string

// CurrentLocation defines a location for the weather forcast.
var CurrentLocation string

// Forecast returns weather forscast for given location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
