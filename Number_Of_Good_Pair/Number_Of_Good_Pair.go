// 1512. Number of Good Pairs
// https://leetcode.com/problems/number-of-good-pairs/description/

/*
Given an array of integers nums, return the number of good pairs.

A pair (i, j) is called good if nums[i] == nums[j] and i < j.
*/

package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	// init result
	total := 0

	// map save (num : num appears)
	countNum := make(map[int]int)

	// loop and record how many times each num appears
	for _, el := range nums {
		countNum[el]++
	}

	// If a number appears n times, then n * (n – 1) // 2 good pairs can be made with this number
	for _, e := range countNum {
		total += e * (e - 1) / 2
	}

	return total
}

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3}))
}
