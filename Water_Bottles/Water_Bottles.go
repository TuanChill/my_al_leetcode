package main

// 1518. Water Bottles
// https://leetcode.com/problems/water-bottles/description/
/* There are numBottles water bottles that are initially full of water. You can exchange numExchange empty water bottles from the market with one full water bottle.

The operation of drinking a full water bottle turns it into an empty bottle.

Given the two integers numBottles and numExchange, return the maximum number of water bottles you can drink.
*/

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	// init first drank
	numBottlesDrink := numBottles

	// loop from drank 2nd
	for numBottles >= numExchange {
		// add numBottles is drunk into numBottlesDrink
		numBottlesDrink += numBottles / numExchange

		// save empty bottles
		numBottles = numBottles/numExchange + numBottles%numExchange

		println("numBottles  ", numBottles)
		println("numBottles Drink  ", numBottlesDrink)
	}
	return numBottlesDrink
}

func main() {
	fmt.Println(numWaterBottles(9, 3))
}
