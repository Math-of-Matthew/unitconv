package unitconv

//Mateus R. Moreira 03/10/2019
//temperature.go is the source of all temperature formulas

//FahrenhetoCelsius converts from Fahrenheit to Celcius
func FahrenhetoCelsius(fahrenheit float32) float32 {
	celsius := (fahrenheit - 32) * 5 / 9
	return celsius
}

// FahrenhetoKelvin converts from Fahrenheit to Kelvin
func FahrenhetoKelvin(fahrenheit float32) float32 {
	kelvin := (fahrenheit + 459.69) * (5 / 9)
	return kelvin
}

// CelsiustoFahrenhe converts from Celcius to Fahrenheit
func CelsiustoFahrenhe(celsius float32) float32 {
	fahrenheit := (celsius * (5 / 9)) + 32
	return fahrenheit
}

// CelsiustoKelvin  converts from Celsius to Kelvin
func CelsiustoKelvin(celsius float32) float32 {
	kelvin := celsius + 273.15
	return kelvin
}

// KelvintoFahrenhe converts from Kelvin to Fahrenheit
func KelvintoFahrenhe(kelvin float32) float32 {
	fahrenheit := (kelvin * (9 / 5)) - 459.67
	return fahrenheit
}

// KelvintoCelsius converts from Kelvin to Celciuss
func KelvintoCelsius(kelvin float32) float32 {
	celsius := kelvin - 273.15
	return celsius
}
