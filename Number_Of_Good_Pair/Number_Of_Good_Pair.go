// 1512. Number of Good Pairs
// https://leetcode.com/problems/number-of-good-pairs/description/

/*
Given an array of integers nums, return the number of good pairs.

A pair (i, j) is called good if nums[i] == nums[j] and i < j.
*/

package main

import "fmt"

// solution 1
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

// solution 2

/*
1 -> 0 good
2 -> 1 good
3 -> 3 good
4 -> 6 good
5 -> 10 good

*/

func numIdenticalPairs2(nums []int) int {
	//init result
	total := 0
	// map to save num: number of good pair
	countGoodPair := make(map[int]int)

	for _, v := range nums {
		if num, ok := countGoodPair[v]; ok {
			// total = current total + num of this number appears
			total += num
			countGoodPair[v]++
		} else {
			// set first value num found
			countGoodPair[v] = 1
		}
	}

	return total
}

func main() {
	fmt.Println(numIdenticalPairs2([]int{1, 2, 3, 1, 1, 3, 1}))
}
