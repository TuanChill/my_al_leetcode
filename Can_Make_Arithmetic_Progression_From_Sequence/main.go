package main

import "fmt"

// 1502. Can Make Arithmetic Progression From Sequence
// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/

/*
A sequence of numbers is called an arithmetic progression if the difference between any two consecutive elements is the same.

Given an array of numbers arr, return true if the array can be rearranged to form an arithmetic progression. Otherwise, return false.



Example 1:

Input: arr = [3,5,1]
Output: true
Explanation: We can reorder the elements as [1,3,5] or [5,3,1] with differences 2 and -2 respectively, between each consecutive elements.
Example 2:

Input: arr = [1,2,4]
Output: false
Explanation: There is no way to reorder the elements to obtain an arithmetic progression.

*/
/*
min, max
min = 1
max = 5

diff 1 = 1 , 2 , 3 , 4 , 5  | len = 5
diff 2 = 1 , 3 , 5 | len = 3

-> diff = (max - min) / (n-1)

(max -min ) % (n -1) -> false

*/

func canMakeArithmeticProgression(arr []int) bool {
	N := len(arr)
	min := arr[0]
	max := arr[0]

	elementInArr := make(map[int]bool)

	for _, v := range arr {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}

		elementInArr[v] = true
	}

	if (max-min)%(N-1) != 0 {
		return false
	}

	diff := (max - min) / (N - 1)
	for min < max {
		min += diff

		if _, isExist := elementInArr[min]; !isExist {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4}))
}
