package main

import (
	"testing"
)

func TestCalculateAvailableCoffees(t *testing.T) {
	ingredients := NewIngredients(1000, 500, 50)
	expectedCoffees := 3

	availableCoffees := calculateAvailableCoffees(ingredients)

	if availableCoffees != expectedCoffees {
		t.Errorf("FAILED expected: %d, given %d", expectedCoffees, availableCoffees)
	} else {
		t.Logf("SUCCESS expected: %d, given %d", expectedCoffees, availableCoffees)
	}
}
