// Package weather ...
package weather

var (
    // CurrentCondition current weather condition.
	CurrentCondition string
    // CurrentLocation current location.
	CurrentLocation  string
)

// Forecast check location weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
