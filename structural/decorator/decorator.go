package decorator

import "errors"

// IngredientAdd is an interface which add ingredient to pizza
type IngredientAdd interface {
	AddIngredient() (string, error)
}

// PizzaDecorator is our decorator which will be added some ingredients
type PizzaDecorator struct {
	Ingredient IngredientAdd
}

// AddIngredient adds pizza ingredient
func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "", errors.New("Not implemented yet")
}

// Meat is an struct which represents a Meat ingredient
type Meat struct {
	Ingredient IngredientAdd
}

// AddIngredient adds a meat ingredient
func (m *Meat) AddIngredient() (string, error) {
	return "", errors.New("Not implemented yet")
}

// Onion is an struct which represents an Onion ingredient
type Onion struct {
	Ingredient IngredientAdd
}

// AddIngredient adds an onion ingredient
func (o *Onion) AddIngredient() (string, error) {
	return "", errors.New("Not implemented yet")
}
