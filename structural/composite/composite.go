package composite

import "fmt"

// Athlete represents an Athlete
type Athlete struct{}

// Train prints a training information
func (a *Athlete) Train() {
	fmt.Println("Training!")
}

// SwimmerA represents an Athlete which swims
type SwimmerA struct {
	MyAthlete Athlete
	MySwim    *func()
}

// Swim prints a swimming information
func Swim() {
	fmt.Println("Swimming!")
}

// Animal represents an Animal
type Animal struct{}

// Eat informs that the animal is eating
func (r *Animal) Eat() {
	fmt.Println("Eating")
}

// Shark represents a animal which Eat and swims
type Shark struct {
	Animal
	Swim func()
}

// Swimmer is an interface that allows an object swim
type Swimmer interface {
	Swim()
}

// Trainer is an interface that wraps the method train
type Trainer interface {
	Train()
}

// SwimmerImpl is the implementation of Swimmer interface
type SwimmerImpl struct{}

// Swim prints a swimming information
func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimming!")
}

// SwimmerB holds Trainer and Swimmer interfaces
type SwimmerB struct {
	Trainer
	Swimmer
}

// Tree represents a binary tree
type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

// Parent is a parent of a son
type Parent struct {
	SomeField int
}

// Son belongs to a parent
type Son struct {
	P Parent
}

// GetParentField returns the parent's field
func GetParentField(p Parent) int {
	return p.SomeField
}
