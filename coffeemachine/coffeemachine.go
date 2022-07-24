package coffeemachine

import (
	"fmt"
	"log"
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
	return Supplies{
		Water:       water,
		Milk:        milk,
		CoffeeBeans: coffeeBeans,
		Cups:        cups,
		Money:       money}
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
	Latte      = "2"
	Cappuccino = "3"
	Back       = "Back"
)

type Coffee struct {
	water       int
	milk        int
	coffeeBeans int
	cost        int
}

func NewCoffee(water, milk, coffeeBeans, cost int) *Coffee {
	return &Coffee{
		water:       water,
		milk:        milk,
		coffeeBeans: coffeeBeans,
		cost:        cost,
	}
}

func getMissingSupply(supplies Supplies, coffee *Coffee) string {
	supplies.Water -= coffee.water
	if supplies.Water < 0 {
		return "water"
	}
	supplies.Milk -= coffee.milk
	if supplies.Milk < 0 {
		return "milk"
	}
	supplies.CoffeeBeans -= coffee.coffeeBeans
	if supplies.CoffeeBeans < 0 {
		return "coffee beans"
	}
	supplies.Cups -= 1
	if supplies.Cups < 0 {
		return "cups"
	}
	return ""
}

func MakeCoffee(supplies *Supplies, coffee *Coffee) {
	if missingSupply := getMissingSupply(*supplies, coffee); missingSupply != "" {
		log.Printf("Sorry, not enough %s!\n", missingSupply)
		return
	}
	log.Println("I have enough resources, making you a coffee!")
	supplies.Water -= coffee.water
	supplies.Milk -= coffee.milk
	supplies.CoffeeBeans -= coffee.coffeeBeans
	supplies.Money += coffee.cost
	supplies.Cups -= 1
	return
}

func coffeeMakerDispatcher(supplies *Supplies, option string) {
	switch option {
	case Espresso:
		coffee := NewCoffee(250, 0, 16, 4)
		MakeCoffee(supplies, coffee)
	case Latte:
		coffee := NewCoffee(350, 75, 20, 7)
		MakeCoffee(supplies, coffee)
	case Cappuccino:
		coffee := NewCoffee(200, 100, 12, 6)
		MakeCoffee(supplies, coffee)
	}

}

func FillCoffeeMachineSupplies(supplies *Supplies, water int, milk int, coffeeBeans int, cups int) {
	supplies.Water += water
	supplies.Milk += milk
	supplies.CoffeeBeans += coffeeBeans
	supplies.Cups += cups
}

func TakeMoneyFromCoffeeMachine(supplies *Supplies, money int) {
	if supplies.Money >= money {
		supplies.Money -= money
	}
}

const (
	buy       = "buy"
	fill      = "fill"
	take      = "take"
	remaining = "remaining"
	exit      = "exit"
)

func makeCoffeeAction(supplies *Supplies) {
	log.Println("What do you want to buy? 1 - espresso, 2 - Latte, 3 - cappuccino, Back - to main menu:")
	var coffeeOption string
	fmt.Scanf("%s", &coffeeOption)
	if coffeeOption == Back {
		return
	}
	coffeeMakerDispatcher(supplies, coffeeOption)
}

func fillCoffeeMachineSuppliesAction(supplies *Supplies) {
	var water, milk, coffeeBeans, cups int
	log.Println("Write how many ml of water you want to add:")
	fmt.Scan(&water)
	log.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&milk)
	log.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&coffeeBeans)
	log.Println("Write how many disposable cups of coffee you want to add:")
	fmt.Scan(&cups)
	FillCoffeeMachineSupplies(supplies, water, milk, coffeeBeans, cups)
}

func takeMoneyFromCoffeeMachineAction(supplies *Supplies) {
	var money = supplies.Money
	TakeMoneyFromCoffeeMachine(supplies, supplies.Money)
	log.Printf("I gave you $%d\n", money)
}

func ActionDispatcher(supplies *Supplies) int {
	log.Println("Write action (buy, fill, take, remaining, exit):")
	var action string
	fmt.Scan(&action)

	switch action {
	case buy:
		makeCoffeeAction(supplies)
	case fill:
		fillCoffeeMachineSuppliesAction(supplies)
	case take:
		takeMoneyFromCoffeeMachineAction(supplies)
	case remaining:
		log.Println(GetCoffeeMachineState(*supplies))
	case exit:
		return 1
	}
	return 0
}
