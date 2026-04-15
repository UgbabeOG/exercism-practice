//Package weather is a package that checks the weather.
package weather

var (
    // CurrentCondition ... this is a global variable.
	CurrentCondition string
    // CurrentLocation ... this is a global variable.
	CurrentLocation  string
)

// Forecast() is a function that takes two string parameters and returns a string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
