package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateAvailableCoffees(t *testing.T) {
	ingredients := NewIngredients(1000, 500, 50)
	expectedCoffees := 3

	availableCoffees := calculateAvailableCoffees(ingredients)

	assert.Equal(t, expectedCoffees, availableCoffees, "Should be equal.")
}
