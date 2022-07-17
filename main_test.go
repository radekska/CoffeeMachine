package main

import (
	"testing"
)

func TestIngredientCalculator(t *testing.T) {
	var expectedWater, expectedMilk, expectedCoffeeBeans = 46800, 11700, 3510
	cups := 234
	var water, milk, coffeeBeans = ingredientCalculator(cups)

	if expectedWater != water || expectedMilk != milk || expectedCoffeeBeans != coffeeBeans {
		t.Errorf("FAILED expected: %d %d %d, given: %d %d %d", expectedWater,
			expectedMilk, expectedCoffeeBeans, water, milk, coffeeBeans)
	} else {
		t.Log("SUCCESS")
	}
}
