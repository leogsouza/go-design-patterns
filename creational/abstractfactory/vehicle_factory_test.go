package abstractfactory

import (
	"testing"
)

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeF.Build(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels\n",
		motorbikeVehicle.NumWheels())

	t.Logf("Motorbike vehicle has %d seats\n",
		motorbikeVehicle.NumSeats())

	sportBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Sport motorbike has type %d\n", sportBike.GetMotorbikeType())

	motorbikeVehicle, err = motorbikeF.Build(CruiseMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("CruiseMotorbike vehicle has %d wheels\n",
		motorbikeVehicle.NumWheels())

	t.Logf("CruiseMotorbike vehicle has %d seats\n",
		motorbikeVehicle.NumSeats())

	cruiseBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Cruise motorbike has type %d\n", cruiseBike.GetMotorbikeType())

	_, err = motorbikeF.Build(20)
	if err == nil {
		t.Error("A motorbike type with ID 20 must return an error")
	}
	t.Log("LOG:", err)
}

func TestCarFactory(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carF.Build(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d wheels\n", carVehicle.NumWheels())
	t.Logf("Car vehicle has %d seats\n", carVehicle.NumSeats())

	luxuryCar, ok := carVehicle.(Car)

	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury car has %d doors.\n", luxuryCar.NumDoors())

	carVehicle, err = carF.Build(FamilyCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d wheels\n", carVehicle.NumWheels())
	t.Logf("Car vehicle has %d seats\n", carVehicle.NumSeats())

	familyCar, ok := carVehicle.(Car)

	if !ok {
		t.Fatal("Struct assertion has falied")
	}

	t.Logf("Family car has %d doors", familyCar.NumDoors())

	_, err = carF.Build(20)
	if err == nil {
		t.Error("A car type with ID 20 must return an error")
	}
	t.Log("LOG:", err)
}

func TestCarNonExistent(t *testing.T) {
	_, err := BuildFactory(10)

	if err == nil {
		t.Error("A Factory type with ID 10 must return an error")
	}
	t.Log("Log:", err)
}
