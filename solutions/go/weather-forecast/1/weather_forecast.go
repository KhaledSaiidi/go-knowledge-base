// Package weather provides tools for forecasting weather.
package weather

var (
	// CurrentCondition describes the current weather condition.
	CurrentCondition string
	// CurrentLocation describes the current location.
	CurrentLocation string
)

// Forecast returns the weather forecast for a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}