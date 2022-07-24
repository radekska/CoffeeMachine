package main

import (
	"coffeemachine/pkg/coffeemachine"
	"log"
)

func main() {
	log.SetFlags(0)
	coffeeMachineSupplies := coffeemachine.NewCoffeeMachineSupplies(400, 540, 120, 9, 550)
	for true {
		state := coffeemachine.ActionDispatcher(&coffeeMachineSupplies)
		if state == 1 {
			break
		}
	}
}
