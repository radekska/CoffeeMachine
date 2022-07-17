package main

import "fmt"

type Ingredients struct {
	water       int
	milk        int
	coffeeBeans int
}

const (
	waterPerCoffee = 200
	milkPerCoffee  = 50
	beansPerCoffee = 15
)

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

func getInput(value int) int {
	_, err := fmt.Scan(&value)
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func main() {
	var neededCups, currentWater, currentMilk, currentBeans int
	fmt.Print("Write how many ml of water the coffee machine has:\n> ")
	currentWater = getInput(currentWater)

	fmt.Print("Write how many ml of milk the coffee machine has:\n> ")
	currentMilk = getInput(currentMilk)

	fmt.Print("Write how many grams of coffee beans the coffee machine has:\n> ")
	currentBeans = getInput(currentBeans)

	fmt.Print("Write how many cups of coffee you will need:\n> ")
	neededCups = getInput(neededCups)

	var currentIngredients = NewIngredients(currentWater, currentMilk, currentBeans)
	var availableCups = calculateAvailableCoffees(currentIngredients)

	if availableCups == neededCups {
		fmt.Println("Yes, I can make that amount of coffee")
	}
	if availableCups < neededCups {
		fmt.Printf("No, I can make only %d cups of coffee\n", availableCups)
	}
	if availableCups > neededCups {
		extraCups := availableCups - neededCups
		fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)\n", extraCups)
	}
}
