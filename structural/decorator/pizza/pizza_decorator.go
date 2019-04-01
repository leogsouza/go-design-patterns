package decorator

import (
	"errors"
	"fmt"
)

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
	return "Pizza with the following ingredients:", nil
}

// Meat is an struct which represents a Meat ingredient
type Meat struct {
	Ingredient IngredientAdd
}

// AddIngredient adds a meat ingredient
func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient" +
			" field of the Meat")
	}
	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "meat"), nil
}

// Onion is an struct which represents an Onion ingredient
type Onion struct {
	Ingredient IngredientAdd
}

// AddIngredient adds an onion ingredient
func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient" +
			" field of the Onion")
	}
	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", s, "onion"), nil
}
