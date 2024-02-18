package main

import (
	"fmt"
)

// 1528. Shuffle String
// https://leetcode.com/problems/shuffle-string/description/

/*
You are given a string s and an integer array indices of the same length. The string s will be shuffled such that the character at the ith position moves to indices[i] in the shuffled string.

Return the shuffled string.

Example 1:


Input: s = "codeleet", indices = [4,5,6,7,0,2,1,3]
Output: "leetcode"
Explanation: As shown, "codeleet" becomes "leetcode" after shuffling.
Example 2:

Input: s = "abc", indices = [0,1,2]
Output: "abc"
Explanation: After shuffling, each character remains in its position.


Constraints:

s.length == indices.length == n
1 <= n <= 100
s consists of only lowercase English letters.
0 <= indices[i] < n
All values of indices are unique.

*/

// solution 1
// 6ms
func restoreString(s string, indices []int) string {
	strArr := make([]string, len(s))

	for i, e := range indices {
		strArr[e] = string(s[i])
	}

	return fmt.Sprint(strArr)
}

// solution 2
// 1ms
func restoreString2(s string, indices []int) string {
	strArr := make([]byte, len(s))

	for i, e := range indices {
		strArr[e] = s[i]
	}

	return string(strArr)
}

func main() {
	fmt.Println(restoreString2("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
}
