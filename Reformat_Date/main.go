package main

import (
	"fmt"
	"strings"
)

// 1507. Reformat Date
// https://leetcode.com/problems/reformat-date/description/

/*
Given a date string in the form Day Month Year, where:

Day is in the set {"1st", "2nd", "3rd", "4th", ..., "30th", "31st"}.
Month is in the set {"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}.
Year is in the range [1900, 2100].
Convert the date string to the format YYYY-MM-DD, where:

YYYY denotes the 4 digit year.
MM denotes the 2 digit month.
DD denotes the 2 digit day.


Example 1:

Input: date = "20th Oct 2052"
Output: "2052-10-20"
Example 2:

Input: date = "6th Jun 1933"
Output: "1933-06-06"
Example 3:

Input: date = "26th May 1960"
Output: "1960-05-26"


Constraints:

The given dates are guaranteed to be valid, so no error handling is necessary.

*/

// solution 1
func reformatDate(date string) string {
	Day := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	var result = make([]string, 3)
	arr := strings.Split(date, " ")

	result[0] = arr[2]

	for i, v := range Day {
		if v == arr[1] {
			if i >= 9 {
				result[1] = fmt.Sprint(i + 1)
			} else {
				result[1] = fmt.Sprintf("0%v", i+1)
			}
		}
	}

	day := arr[0][:len(arr[0])-2]

	if len(day) == 2 {
		result[2] = day
	} else {
		result[2] = fmt.Sprintf("0%v", day)
	}

	return strings.Join(result, "-")
}

// solution 2

func reformatDate2(date string) string {
	result := ""
	Month := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}

	if len(date) == 12 {
		date = "0" + date
	}

	year := date[9:]
	month := Month[date[5:8]]
	day := date[:2]

	result = year + "-" + month + "-" + day

	return result
}

func main() {
	fmt.Println(reformatDate2("20th Sep 2052"))
}
