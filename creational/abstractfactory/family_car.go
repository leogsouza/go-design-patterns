package abstractfactory

// FamilyCar represents a car of type familiar
type FamilyCar struct{}

// NumDoors implements Car interface
func (*FamilyCar) NumDoors() int {
	return 4
}

// NumWheels implements Vehicle interface
func (*FamilyCar) NumWheels() int {
	return 4
}

// NumSeats implements Vehicle interface
func (*FamilyCar) NumSeats() int {
	return 5
}
