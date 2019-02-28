package abstractfactory

// SportMotorbike is a type which represents a Super Motobike
type SportMotorbike struct{}

// NumWheels returns how many wheels the motobike has
func (s *SportMotorbike) NumWheels() int {
	return 2
}

// NumSeats returns how many seats the motorbike has
func (s *SportMotorbike) NumSeats() int {
	return 1
}

// GetMotorbikeType returns the type of motorbike according the type
func (s *SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}
