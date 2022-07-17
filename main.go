package coffeeMachine

import "fmt"

func ingredientCalculator(cups int) (int, int, int) {
	var water, milk, coffeeBeans = cups * 200, cups * 50, cups * 15
	return water, milk, coffeeBeans
}

func coffeeMachine() {
	fmt.Print("Write how many cups of coffee you will need:\n> ")
	var cups int
	_, err := fmt.Scan(&cups)
	if err != nil {
		fmt.Println(err)
	}
	var water, milk, coffeeBeans = ingredientCalculator(cups)
	fmt.Printf("For %d cups of coffee you will need:\n", cups)
	fmt.Printf("%d ml of water\n", water)
	fmt.Printf("%d ml of milk\n", milk)
	fmt.Printf("%d g of coffe coffeeBeans\n", coffeeBeans)
}
