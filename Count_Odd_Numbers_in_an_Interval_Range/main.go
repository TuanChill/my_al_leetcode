package main

import "fmt"

// 1523. Count Odd Numbers in an Interval Range
// https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/description/

/*
Given two non-negative integers low and high. Return the count of odd numbers between low and high (inclusive).



Example 1:

Input: low = 3, high = 7
Output: 3
Explanation: The odd numbers between 3 and 7 are [3,5,7].
Example 2:

Input: low = 8, high = 10
Output: 1
Explanation: The odd numbers between 8 and 10 are [9].

*/

// 3 -5

// 3 ,5
// 2, 4

// 8- 10
// 9
// 8, 10

// 3- 6

// 4, 6
// 3, 5

func countOdds(low int, high int) int {
	if low%2 == 1 && high%2 == 1 {
		return (high-low+1)/2 + 1
	}
	return (high - low + 1) / 2
}

func main() {
	fmt.Println(countOdds(3, 9))
}
