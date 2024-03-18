package main

import "fmt"

// 1480. Running Sum of 1d Array

// https://leetcode.com/problems/running-sum-of-1d-array/description/

/*

Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

Return the running sum of nums.



Example 1:

Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
Example 2:

Input: nums = [1,1,1,1,1]
Output: [1,2,3,4,5]
Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].
Example 3:

Input: nums = [3,1,2,10,1]
Output: [3,4,6,16,17]


Constraints:

1 <= nums.length <= 1000
-10^6 <= nums[i] <= 10^6

*/

func runningSum(nums []int) []int {
	result := []int{}

	for i, v := range nums {
		if len(result) == 0 {
			result = append(result, v)
		} else {
			result = append(result, v+result[i-1])
		}
	}

	return result
}

func runningSum2(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return nums
}

func main() {
	fmt.Println(runningSum2([]int{1, 2, 3, 4}))
}
