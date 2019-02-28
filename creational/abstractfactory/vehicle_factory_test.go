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

	sportBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Sport motorbike has tpe %d\n", sportBike.GetMotorbikeType())
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

	familyCar, ok := carVehicle.(Car)

	if !ok {
		t.Fatal("Struct assertion has falied")
	}

	t.Logf("Family car has %d doors", familyCar.NumDoors())
	t.Logf("Family car has %d seats\n", carVehicle.NumSeats())

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
