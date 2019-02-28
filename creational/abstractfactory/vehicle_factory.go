package abstractfactory

import "fmt"

// VehicleFactory is the interface by an object that creates
// a new Vehicle
type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

// Factory types
const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

// BuildFactory create a new factory according the type
func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	default:
		return nil, fmt.Errorf("factory with id %d not recognized", f)
	}
}
