package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct{}

func (f *ManufacturingDirector) Construct() {
	// Implementation goes here
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	// Implementation goes here
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct{}

func (c *CarBuilder) SetWheels() BuildProcess {
	return nil
}

func (c *CarBuilder) SetSeats() BuildProcess {
	return nil
}

func (c *CarBuilder) SetStructure() BuildProcess {
	return nil
}

func (c *CarBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}
