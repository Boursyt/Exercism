// Package weather provides basic weather forecasting functionality.
package weather

var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string
	// CurrentLocation represents the current location (city) for the forecast.
	CurrentLocation string
)

// Forecast updates the current location and condition and returns the formatted forecast string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
