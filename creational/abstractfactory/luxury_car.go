package abstractfactory

// LuxuryCar represents a car of type familiar
type LuxuryCar struct{}

// NumDoors implements Car interface
func (*LuxuryCar) NumDoors() int {
	return 4
}

// NumWheels implements Vehicle interface
func (*LuxuryCar) NumWheels() int {
	return 4
}

// NumSeats implements Vehicle interface
func (*LuxuryCar) NumSeats() int {
	return 5
}
