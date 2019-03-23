package composite

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training!")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    *func()
}

func Swim() {
	fmt.Println("Swimming!")
}

type Animal struct{}

func (r *Animal) Eat() {
	fmt.Println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

type Parent struct {
	SomeField int
}

type Son struct {
	P Parent
}

func GetParentField(p Parent) int {
	return p.SomeField
}
