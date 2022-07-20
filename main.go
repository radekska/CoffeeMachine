package main

import (
	"fmt"
	"strings"
)

const (
	waterPerCoffee = 200
	milkPerCoffee  = 50
	beansPerCoffee = 15
)

type Ingredients struct {
	water       int
	milk        int
	coffeeBeans int
}

type CoffeeMachineSupplies struct {
	ingredients Ingredients
	cups        int
	money       int
}

func NewCoffeeMachineSupplies(ingredients Ingredients, cups int, money int) CoffeeMachineSupplies {
	return CoffeeMachineSupplies{ingredients: ingredients, cups: cups, money: money}
}

func NewIngredients(water int, milk int, coffeeBeans int) Ingredients {
	return Ingredients{
		water:       water,
		milk:        milk,
		coffeeBeans: coffeeBeans,
	}
}

var singleCoffeeIngredients = NewIngredients(waterPerCoffee, milkPerCoffee, beansPerCoffee)

func calculateAvailableCoffees(currentIngredients Ingredients) int {
	availableWater := currentIngredients.water / singleCoffeeIngredients.water
	availableMilk := currentIngredients.milk / singleCoffeeIngredients.milk
	availableCoffeeBeans := currentIngredients.coffeeBeans / singleCoffeeIngredients.coffeeBeans
	availableIngredients := [3]int{availableWater, availableMilk, availableCoffeeBeans}

	minimal := availableIngredients[0]
	for _, availableIngredient := range availableIngredients {
		if availableIngredient < minimal {
			minimal = availableIngredient
		}
	}
	return minimal
}

func getInput(value int) {
	_, err := fmt.Scan(&value)
	if err != nil {
		fmt.Println(err)
	}
}

func getCoffeeMachineState(coffeeMachineSupplies CoffeeMachineSupplies) string {
	var messages = []string{"The coffee machine has:", "%d ml of water", "%d ml of milk", "%d g of coffee beans",
		"%d disposable cups", "$%d of money"}
	message := strings.Join(messages, "\n")
	return fmt.Sprintf(message, coffeeMachineSupplies.ingredients.water, coffeeMachineSupplies.ingredients.milk,
		coffeeMachineSupplies.ingredients.coffeeBeans, coffeeMachineSupplies.cups, coffeeMachineSupplies.money)
}

const (
	espresso   = iota
	latte      = iota
	cappuccino = iota
)

func makeCoffee(coffeeMachineSupplies *CoffeeMachineSupplies, option int) {
	switch option {
	case espresso:
		coffeeMachineSupplies.ingredients.water -= 250
		coffeeMachineSupplies.ingredients.coffeeBeans -= 16
		coffeeMachineSupplies.money += 4
	case latte:
		coffeeMachineSupplies.ingredients.water -= 350
		coffeeMachineSupplies.ingredients.milk -= 75
		coffeeMachineSupplies.ingredients.coffeeBeans -= 20
		coffeeMachineSupplies.money += 7
	case cappuccino:
		coffeeMachineSupplies.ingredients.water -= 200
		coffeeMachineSupplies.ingredients.milk -= 100
		coffeeMachineSupplies.ingredients.coffeeBeans -= 12
		coffeeMachineSupplies.money += 6
	}
	if option == espresso || option == latte || option == cappuccino {
		coffeeMachineSupplies.cups -= 1
	}

}

func fillCoffeeMachineSupplies(coffeeMachineSupplies *CoffeeMachineSupplies, water int, milk int, coffeeBeans int, cups int) {
	coffeeMachineSupplies.ingredients.water += water
	coffeeMachineSupplies.ingredients.milk += milk
	coffeeMachineSupplies.ingredients.coffeeBeans += coffeeBeans
	coffeeMachineSupplies.cups += cups
}

func takeMoneyFromCoffeeMachine(coffeeMachineSupplies *CoffeeMachineSupplies, money int) {
	if coffeeMachineSupplies.money >= money {
		coffeeMachineSupplies.money -= money
	}
}

const (
	buy  = "buy"
	fill = "fill"
	take = "take"
)

func writeAction(coffeeMachineSupplies *CoffeeMachineSupplies) {
	fmt.Println("Write action (buy, fill, take):")
	var action string
	fmt.Scanf("%s", &action)

	switch action {
	case buy:
		fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
		var coffee int
		fmt.Scanf("%d", &coffee)
		makeCoffee(coffeeMachineSupplies, coffee)
	}
}

func main() {
	coffeeMachineSupplies := NewCoffeeMachineSupplies(NewIngredients(400, 540, 120), 9, 550)
	fmt.Println(getCoffeeMachineState(coffeeMachineSupplies))
	writeAction(&coffeeMachineSupplies)
	fmt.Println(getCoffeeMachineState(coffeeMachineSupplies))
}
