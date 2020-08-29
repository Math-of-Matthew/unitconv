// The length provide you a source of unit
// convetion of lenght

package unitconv

const (
	// Kmmi is a value used to convert Kilometers to Miles
	Kmmi = 0.62137
	// Yame is a value used to convert Yards and Miles
	Yame = 1.0936
	// Ftya is used to convert feet to yards
	Ftya = 0.33333
)

// MilestoKilometers returns Kilometers in x Miles
func MilestoKilometers(kilometers float32) float32 {
	return kilometers * Kmmi

}

// KilometerstoMiles return Miles in x Kilometers
func KilometerstoMiles(miles float32) float32 {
	return miles / Kmmi
}

// FeetoMeters converts feet to meters
func FeetoMeters(feet float32) float32 {
	feet *= 0.3048
	return feet
}

// FeetoYard (feet float32) returns flaot32 in yard value
func FeetoYard(feet float32) float32 {
	return feet * 0.33333
}

// MeterstoFeet converts Meters to feet
func MeterstoFeet(meter float32) float32 {
	return meter * 3.2808
}

// InchestoCentimeter (inches float32) returns float32 in Centimeter
func InchestoCentimeter(inches float32) float32 {
	return inches / 0.39370
}

// InchestoMillimeters (inches float32) returns float32 in Millimeters value
func InchestoMillimeters(inches float32) float32 {
	return InchestoCentimeter(inches) * 10
}

// InchestoFeet (inches float32) returns float32 in Feet value
func InchestoFeet(inches float32) float32 {
	return inches * 0.083333
}

// MeterstoYard (meters float32) returns float32 in Yard value
func MeterstoYard(meters float32) float32 {
	return meters * Yame
}

// YardtoMeters (yard float32) returns float32 in Meters value
func YardtoMeters(yard float32) float32 {
	return yard / Yame
}

// YardtoFeet (feet float32) returns float32 in Yard value
func YardtoFeet(yard float32) float32 {
	return yard / 3.0000
}
