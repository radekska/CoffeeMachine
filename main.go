package main

import (
	"fmt"
	"strings"
)

type CoffeeMachineSupplies struct {
	water       int
	milk        int
	coffeeBeans int
	cups        int
	money       int
}

func NewCoffeeMachineSupplies(water int, milk int, coffeeBeans int, cups int, money int) CoffeeMachineSupplies {
	return CoffeeMachineSupplies{water: water, milk: milk, coffeeBeans: coffeeBeans, cups: cups, money: money}
}

func getCoffeeMachineState(coffeeMachineSupplies CoffeeMachineSupplies) string {
	var messages = []string{"The coffee machine has:", "%d ml of water", "%d ml of milk", "%d g of coffee beans",
		"%d disposable cups", "$%d of money"}
	message := strings.Join(messages, "\n")
	return fmt.Sprintf(message, coffeeMachineSupplies.water, coffeeMachineSupplies.milk,
		coffeeMachineSupplies.coffeeBeans, coffeeMachineSupplies.cups, coffeeMachineSupplies.money)
}

const (
	espresso   = "1"
	latte      = "2"
	cappuccino = "3"
	back       = "back"
)

func makeCoffee(coffeeMachineSupplies *CoffeeMachineSupplies, option string) {
	switch option {
	case espresso:
		coffeeMachineSupplies.water -= 250
		coffeeMachineSupplies.coffeeBeans -= 16
		coffeeMachineSupplies.money += 4
	case latte:
		coffeeMachineSupplies.water -= 350
		coffeeMachineSupplies.milk -= 75
		coffeeMachineSupplies.coffeeBeans -= 20
		coffeeMachineSupplies.money += 7
	case cappuccino:
		coffeeMachineSupplies.water -= 200
		coffeeMachineSupplies.milk -= 100
		coffeeMachineSupplies.coffeeBeans -= 12
		coffeeMachineSupplies.money += 6
	}
	if option == espresso || option == latte || option == cappuccino {
		coffeeMachineSupplies.cups -= 1
	}
}

func fillCoffeeMachineSupplies(coffeeMachineSupplies *CoffeeMachineSupplies, water int, milk int, coffeeBeans int, cups int) {
	coffeeMachineSupplies.water += water
	coffeeMachineSupplies.milk += milk
	coffeeMachineSupplies.coffeeBeans += coffeeBeans
	coffeeMachineSupplies.cups += cups
}

func takeMoneyFromCoffeeMachine(coffeeMachineSupplies *CoffeeMachineSupplies, money int) {
	if coffeeMachineSupplies.money >= money {
		coffeeMachineSupplies.money -= money
	}
}

const (
	buy       = "buy"
	fill      = "fill"
	take      = "take"
	remaining = "remaining"
	exit      = "exit"
)

func makeCoffeeAction(coffeeMachineSupplies *CoffeeMachineSupplies) {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	var coffeeOption string
	fmt.Scanf("%s", &coffeeOption)
	if coffeeOption == back {
		return
	}
	makeCoffee(coffeeMachineSupplies, coffeeOption)
}

func fillCoffeeMachineSuppliesAction(coffeeMachineSupplies *CoffeeMachineSupplies) {
	var water, milk, coffeeBeans, cups int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&water)
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&milk)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&coffeeBeans)
	fmt.Println("Write how many disposable cups of coffee you want to add:")
	fmt.Scan(&cups)
	fillCoffeeMachineSupplies(coffeeMachineSupplies, water, milk, coffeeBeans, cups)
}

func takeMoneyFromCoffeeMachineAction(coffeeMachineSupplies *CoffeeMachineSupplies) {
	var money = coffeeMachineSupplies.money
	takeMoneyFromCoffeeMachine(coffeeMachineSupplies, coffeeMachineSupplies.money)
	fmt.Printf("I gave you $%d\n", money)
}

func actionDispatcher(coffeeMachineSupplies *CoffeeMachineSupplies) int {
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	var action string
	fmt.Scan(&action)

	switch action {
	case buy:
		makeCoffeeAction(coffeeMachineSupplies)
	case fill:
		fillCoffeeMachineSuppliesAction(coffeeMachineSupplies)
	case take:
		takeMoneyFromCoffeeMachineAction(coffeeMachineSupplies)
	case remaining:
		fmt.Println(getCoffeeMachineState(*coffeeMachineSupplies))
	case exit:
		return 1
	}
	return 0
}

func main() {
	coffeeMachineSupplies := NewCoffeeMachineSupplies(400, 540, 120, 9, 550)
	for true {
		state := actionDispatcher(&coffeeMachineSupplies)
		if state == 1 {
			break
		}
	}
}
