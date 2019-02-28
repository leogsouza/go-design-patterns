package abstractfactory

import (
	"fmt"
)

// Car types
const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

// CarFactory represents the factory of cars
type CarFactory struct{}

// Build instantiates a new Vehicle according type passed
func (c *CarFactory) Build(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized", v)
	}
}
