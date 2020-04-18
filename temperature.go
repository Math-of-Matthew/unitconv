package unitconv

//Mateus R. Moreira 03/10/2019
//temperature.go is the source of all temperature formulas

//FahrenhetoCelcius converts from Fahrenheit to Celcius
func FahrenhetoCelcius(fahrenheit float32) float32 {
	celcius := (fahrenheit - 32) * 5 / 9
	return celcius
}

// FahrenhetoKelvin converts from Fahrenheit to Kelvin
func FahrenhetoKelvin(fahrenheit float32) float32 {
	kelvin := (fahrenheit + 459.69) * (5 / 9)
	return kelvin
}

// CelciustoFahrenhe converts from Celcius to Fahrenheit
func CelciustoFahrenhe(celcius float32) float32 {
	fahrenheit := (celcius * (5 / 9)) + 32
	return fahrenheit
}

// CelciustoKelvin  converts from Celcius to Kelvin
func CelciustoKelvin(celcius float32) float32 {
	kelvin := celcius + 273.15
	return kelvin
}

// KelvintoFahrenhe converts from Kelvin to Fahrenheit
func KelvintoFahrenhe(kelvin float32) float32 {
	fahrenheit := (kelvin * (9 / 5)) - 459.67
	return fahrenheit
}

// KelvintoCelcius converts from Kelvin to Celciuss
func KelvintoCelcius(kelvin float32) float32 {
	celcius := kelvin - 273.15
	return celcius
}
