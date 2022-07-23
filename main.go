package main

import "coffeemachine/coffeemachine"

func main() {
	coffeeMachineSupplies := coffeemachine.NewCoffeeMachineSupplies(400, 540, 120, 9, 550)
	for true {
		state := coffeemachine.ActionDispatcher(&coffeeMachineSupplies)
		if state == 1 {
			break
		}
	}
}
