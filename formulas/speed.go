package unitconv

// KnottoKilomerh (knot float32) float32
func KnottoKilomerh(knot float32) float32 {
	return knot * 1.852
}

// KnottoMeterS (knot float32) float32
func KnottoMeterS(knot float32) float32 {
	return knot / 1.944
}

// KnottoFeetS (knot float32) float32
func KnottoFeetS(knot float32) float32 {
	return knot * 1.688
}

// KnottoMilesh (knot float32) float32
func KnottoMilesh(knot float32) float32 {
	return knot * 1.151
}

// MileshtoKilometerh (milesh float32) float32
func MileshtoKilometerh(milesh float32) float32 {
	return milesh * 1.609
}

// MileshtoMeters (milesh float32) float32
func MileshtoMeters(milesh float32) float32 {
	return milesh / 2.237
}

// MileshtoFoots (milesh float32) float32
func MileshtoFoots(milesh float32) float32 {
	return milesh * 1.467
}

// MileshtoKnot (milesh float32) float32
func MileshtoKnot(milesh float32) float32 {
	return milesh * 1.51
}

// KilometerhtoKnot (Kilometerh float32) float32
func KilometerhtoKnot(Kilometerh float32) float32 {
	return Kilometerh / 1.852
}

// KilometerhtoMeters (Kilometerh float32) float32
func KilometerhtoMeters(Kilometerh float32) float32 {
	return Kilometerh / 3.6
}

// KilometerhtoMilesh (Kilometerh float32) float32
func KilometerhtoMilesh(Kilometerh float32) float32 {
	return Kilometerh / 1.609
}

// KilometerhtoFeets (Kilometerh float32) float32
func KilometerhtoFeets(Kilometerh float32) float32 {
	return Kilometerh / 1.097
}

// MeterstoMilesh (meters float32) float32
func MeterstoMilesh(meters float32) float32 {
	return meters * 2.237
}

// MeterstoKnot (meters float32) float32
func MeterstoKnot(meters float32) float32 {
	return meters * 1.944
}

// MeterstoKilometerh (meters float32) float32
func MeterstoKilometerh(meters float32) float32 {
	return meters * 3.6
}

// MeterstoFoots (meters float32) float32
func MeterstoFoots(meters float32) float32 {
	return meters * 3.281
}

// FootstoKnot (foots float32) float32
func FootstoKnot(foots float32) float32 {
	return foots / 1.688
}

// FootstoKilometerh (foots float32) float32
func FootstoKilometerh(foots float32) float32 {
	return foots * 1.097
}

// FootstoMeters (foots float32) float32
func FootstoMeters(foots float32) float32 {
	return foots / 3.281
}

// FootstoMilesh (foots float32) float32
func FootstoMilesh(foots float32) float32 {
	return foots / 1.467
}
