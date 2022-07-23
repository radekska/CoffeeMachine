package coffeemachine

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestGetCoffeeMachineState(t *testing.T) {
	coffeeMachineSupplies := NewCoffeeMachineSupplies(100, 50, 10, 5, 1000)

	message := GetCoffeeMachineState(coffeeMachineSupplies)

	assert.Equal(t, "The coffee machine has:\n100 ml of water\n50 ml of milk\n10 g of coffee beans\n5 disposable cups\n$1000 of money", message)
}

func TestMakeCoffee(t *testing.T) {
	coffeeMachineSupplies := NewCoffeeMachineSupplies(1000, 100, 50, 5, 1000)

	coffeeMakerDispatcher(&coffeeMachineSupplies, Espresso)

	assert.Equal(t, 750, coffeeMachineSupplies.Water)
	assert.Equal(t, 34, coffeeMachineSupplies.CoffeeBeans)
	assert.Equal(t, 1004, coffeeMachineSupplies.Money)
	assert.Equal(t, 4, coffeeMachineSupplies.Cups)

}

func TestFillCoffeeMachine(t *testing.T) {
	coffeeMachineSupplies := NewCoffeeMachineSupplies(1000, 100, 50, 5, 1000)

	FillCoffeeMachineSupplies(&coffeeMachineSupplies, 500, 250, 10, 30)

	assert.Equal(t, 1500, coffeeMachineSupplies.Water)
	assert.Equal(t, 350, coffeeMachineSupplies.Milk)
	assert.Equal(t, 60, coffeeMachineSupplies.CoffeeBeans)
	assert.Equal(t, 35, coffeeMachineSupplies.Cups)
}

func TestTakeMoneyFromCoffeeMachine(t *testing.T) {
	coffeeMachineSupplies := NewCoffeeMachineSupplies(1000, 100, 50, 5, 1000)

	TakeMoneyFromCoffeeMachine(&coffeeMachineSupplies, 1000)

	assert.Equal(t, 0, coffeeMachineSupplies.Money)
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func TestCoffeeMakerDispatcher(t *testing.T) {
	var tests = []struct {
		supplies       Supplies
		option         string
		expectedOutput string
	}{
		{NewCoffeeMachineSupplies(0, 0, 0, 0, 0), Espresso, "Sorry, not enough water!\n"},
		{NewCoffeeMachineSupplies(300, 0, 0, 0, 0), Espresso, "Sorry, not enough coffee beans!\n"},
		{NewCoffeeMachineSupplies(300, 0, 50, 0, 0), Espresso, "Sorry, not enough cups!\n"},
		{NewCoffeeMachineSupplies(300, 0, 50, 5, 0), Espresso, "I have enough resources, making you a coffee!\n"},
		{NewCoffeeMachineSupplies(0, 0, 0, 0, 0), Latte, "Sorry, not enough water!\n"},
	}
	log.SetFlags(0)
	// TODO - add cases for the rest of coffees
	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%s", tt.supplies, tt.option)
		t.Run(testName, func(t *testing.T) {
			output := captureOutput(func() {
				coffeeMakerDispatcher(&tt.supplies, tt.option)
			})
			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}
