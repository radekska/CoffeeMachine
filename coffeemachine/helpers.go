package coffeemachine

import "fmt"

func GetStringInput() string {
	fmt.Print("> ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}
	return input
}

func GetIntInput() int {
	fmt.Print("> ")
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}
	return input
}
