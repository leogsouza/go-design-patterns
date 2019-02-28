package abstractfactory

// CruiseMotorbike is a type which represents a Cruise motorbike
type CruiseMotorbike struct{}

// NumWheels returns how many wheels the motorbike has
func (c *CruiseMotorbike) NumWheels() int {
	return 2
}

// NumSeats returns how many seats the motorbike has
func (c *CruiseMotorbike) NumSeats() int {
	return 2
}

// GetMotorbikeType returns the type of motorbike
func (c *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}
