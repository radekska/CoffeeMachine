package test

import (
	"coffeemachine/coffeemachine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCoffeeMachineState(t *testing.T) {
	coffeeMachineSupplies := coffeemachine.NewCoffeeMachineSupplies(100, 50, 10, 5, 1000)

	message := coffeemachine.GetCoffeeMachineState(coffeeMachineSupplies)

	assert.Equal(t, "The coffee machine has:\n100 ml of water\n50 ml of milk\n10 g of coffee beans\n5 disposable cups\n$1000 of money", message)
}

func TestMakeCoffee(t *testing.T) {
	coffeeMachineSupplies := coffeemachine.NewCoffeeMachineSupplies(1000, 100, 50, 5, 1000)

	coffeemachine.MakeCoffee(&coffeeMachineSupplies, coffeemachine.Espresso)

	assert.Equal(t, 750, coffeeMachineSupplies.Water)
	assert.Equal(t, 34, coffeeMachineSupplies.CoffeeBeans)
	assert.Equal(t, 1004, coffeeMachineSupplies.Money)
	assert.Equal(t, 4, coffeeMachineSupplies.Cups)

}

func TestFillCoffeeMachine(t *testing.T) {
	coffeeMachineSupplies := coffeemachine.NewCoffeeMachineSupplies(1000, 100, 50, 5, 1000)

	coffeemachine.FillCoffeeMachineSupplies(&coffeeMachineSupplies, 500, 250, 10, 30)

	assert.Equal(t, 1500, coffeeMachineSupplies.Water)
	assert.Equal(t, 350, coffeeMachineSupplies.Milk)
	assert.Equal(t, 60, coffeeMachineSupplies.CoffeeBeans)
	assert.Equal(t, 35, coffeeMachineSupplies.Cups)
}

func TestTakeMoneyFromCoffeeMachine(t *testing.T) {
	coffeeMachineSupplies := coffeemachine.NewCoffeeMachineSupplies(1000, 100, 50, 5, 1000)

	coffeemachine.TakeMoneyFromCoffeeMachine(&coffeeMachineSupplies, 1000)

	assert.Equal(t, 0, coffeeMachineSupplies.Money)
}
