package main

import (
	"fmt"
	"strings"
)

// 58. Length of Last Word
// https://leetcode.com/problems/length-of-last-word/description/

/*
Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal
substring
 consisting of non-space characters only.



Example 1:

Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.
Example 2:

Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.
Example 3:

Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.


Constraints:

1 <= s.length <= 104
s consists of only English letters and spaces ' '.
There will be at least one word in s.
*/

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	count := 0
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else {
			return count
		}
	}

	return count
}

func lengthOfLastWord2(s string) int {
	split := strings.Fields(s)
	return len(split[len(split)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}
