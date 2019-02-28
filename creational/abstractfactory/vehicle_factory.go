package abstractfactory

// VehicleFactory is the interface by an object that creates
// a new Vehicle
type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}
