package abstractfactory

import "fmt"

// Types of motorbike
const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

// MotorbikeFactory is a type which represents a factory of motorbikes
type MotorbikeFactory struct{}

// Build returns a new Motorbike type according the type
func (m *MotorbikeFactory) Build(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized", v)
	}
}
