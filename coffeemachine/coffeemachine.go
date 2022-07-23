package coffeemachine

import (
	"fmt"
	"strings"
)

type Supplies struct {
	Water       int
	Milk        int
	CoffeeBeans int
	Cups        int
	Money       int
}

func NewCoffeeMachineSupplies(water int, milk int, coffeeBeans int, cups int, money int) Supplies {
	return Supplies{Water: water, Milk: milk, CoffeeBeans: coffeeBeans, Cups: cups, Money: money}
}

func GetCoffeeMachineState(coffeeMachineSupplies Supplies) string {
	var messages = []string{"The coffee machine has:", "%d ml of water", "%d ml of milk", "%d g of coffee beans",
		"%d disposable cups", "$%d of money"}
	message := strings.Join(messages, "\n")
	return fmt.Sprintf(message, coffeeMachineSupplies.Water, coffeeMachineSupplies.Milk,
		coffeeMachineSupplies.CoffeeBeans, coffeeMachineSupplies.Cups, coffeeMachineSupplies.Money)
}

const (
	Espresso   = "1"
	latte      = "2"
	cappuccino = "3"
	back       = "back"
)

func MakeCoffee(coffeeMachineSupplies *Supplies, option string) {
	switch option {
	case Espresso:
		coffeeMachineSupplies.Water -= 250
		coffeeMachineSupplies.CoffeeBeans -= 16
		coffeeMachineSupplies.Money += 4
	case latte:
		coffeeMachineSupplies.Water -= 350
		coffeeMachineSupplies.Milk -= 75
		coffeeMachineSupplies.CoffeeBeans -= 20
		coffeeMachineSupplies.Money += 7
	case cappuccino:
		coffeeMachineSupplies.Water -= 200
		coffeeMachineSupplies.Milk -= 100
		coffeeMachineSupplies.CoffeeBeans -= 12
		coffeeMachineSupplies.Money += 6
	}
	if option == Espresso || option == latte || option == cappuccino {
		coffeeMachineSupplies.Cups -= 1
	}
}

func FillCoffeeMachineSupplies(coffeeMachineSupplies *Supplies, water int, milk int, coffeeBeans int, cups int) {
	coffeeMachineSupplies.Water += water
	coffeeMachineSupplies.Milk += milk
	coffeeMachineSupplies.CoffeeBeans += coffeeBeans
	coffeeMachineSupplies.Cups += cups
}

func TakeMoneyFromCoffeeMachine(coffeeMachineSupplies *Supplies, money int) {
	if coffeeMachineSupplies.Money >= money {
		coffeeMachineSupplies.Money -= money
	}
}

const (
	buy       = "buy"
	fill      = "fill"
	take      = "take"
	remaining = "remaining"
	exit      = "exit"
)

func makeCoffeeAction(coffeeMachineSupplies *Supplies) {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	var coffeeOption string
	fmt.Scanf("%s", &coffeeOption)
	if coffeeOption == back {
		return
	}
	MakeCoffee(coffeeMachineSupplies, coffeeOption)
}

func fillCoffeeMachineSuppliesAction(coffeeMachineSupplies *Supplies) {
	var water, milk, coffeeBeans, cups int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&water)
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&milk)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&coffeeBeans)
	fmt.Println("Write how many disposable cups of coffee you want to add:")
	fmt.Scan(&cups)
	FillCoffeeMachineSupplies(coffeeMachineSupplies, water, milk, coffeeBeans, cups)
}

func takeMoneyFromCoffeeMachineAction(coffeeMachineSupplies *Supplies) {
	var money = coffeeMachineSupplies.Money
	TakeMoneyFromCoffeeMachine(coffeeMachineSupplies, coffeeMachineSupplies.Money)
	fmt.Printf("I gave you $%d\n", money)
}

func ActionDispatcher(coffeeMachineSupplies *Supplies) int {
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
		fmt.Println(GetCoffeeMachineState(*coffeeMachineSupplies))
	case exit:
		return 1
	}
	return 0
}
