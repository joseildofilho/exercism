// Package weather constains something related to the weather.
package weather

// CurrentCondition no good way to document a variable.
var CurrentCondition string
// CurrentLocation  no good way to document a variable.
var CurrentLocation string

//Forecast the same should self explanatory.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
