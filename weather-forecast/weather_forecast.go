// Package weather forecasts the current state of the weather for a particular locations.
package weather

// CurrentCondition is the current weather condition as a string.
var CurrentCondition string
// CurrentLocation is the current location as a string.
var CurrentLocation string

// Forecast finction shows the weather for a oarticular location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
