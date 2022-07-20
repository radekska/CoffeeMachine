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

func TestGetCoffeeMachineState(t *testing.T) {
	coffeeMachineSupplies := NewCoffeeMachineSupplies(NewIngredients(100, 50, 10), 5, 1000)

	message := getCoffeeMachineState(coffeeMachineSupplies)

	assert.Equal(t, "The coffee machine has:\n100 ml of water\n50 ml of milk\n10 g of coffee beans\n5 disposable cups\n$1000 of money", message)
}

func TestMakeCoffee(t *testing.T) {
	coffeeMachineSupplies := NewCoffeeMachineSupplies(NewIngredients(1000, 100, 50), 5, 1000)

	makeCoffee(&coffeeMachineSupplies, espresso)

	assert.Equal(t, 750, coffeeMachineSupplies.ingredients.water)
	assert.Equal(t, 34, coffeeMachineSupplies.ingredients.coffeeBeans)
	assert.Equal(t, 1004, coffeeMachineSupplies.money)
	assert.Equal(t, 4, coffeeMachineSupplies.cups)

}

func TestFillCoffeeMachine(t *testing.T) {
	coffeeMachineSupplies := NewCoffeeMachineSupplies(NewIngredients(1000, 100, 50), 5, 1000)

	fillCoffeeMachineSupplies(&coffeeMachineSupplies, 500, 250, 10, 30)

	assert.Equal(t, 1500, coffeeMachineSupplies.ingredients.water)
	assert.Equal(t, 350, coffeeMachineSupplies.ingredients.milk)
	assert.Equal(t, 60, coffeeMachineSupplies.ingredients.coffeeBeans)
	assert.Equal(t, 35, coffeeMachineSupplies.cups)
}

func TestTakeMoneyFromCoffeeMachine(t *testing.T) {
	coffeeMachineSupplies := NewCoffeeMachineSupplies(NewIngredients(1000, 100, 50), 5, 1000)

	takeMoneyFromCoffeeMachine(&coffeeMachineSupplies, 1000)

	assert.Equal(t, 0, coffeeMachineSupplies.money)
}
