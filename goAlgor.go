package main

import (
	"fmt"
)

func main() {
	// test1 := []int{1, 3, 4, 4, 4, 5}
	// fmt.Println(mode(test1))

	test2 := 2000
	test3 := 1900
	test4 := 2016
	isLeapYear(test2)
	isLeapYear(test3)
	isLeapYear(test4)

	test5 := []int{1, 2, 5}
	test6 := []int{1, 2, 3, 4, 5}
	test7 := []int{3, 4, 5}
	fmt.Println(intNotInArr(test5))
	fmt.Println(intNotInArr(test6))
	fmt.Println(intNotInArr(test7))
}

// func mode(arr []int) int {
// 	maxCount := 1
// 	var counts []int
// 	maxNum := arr[0]
// 	for i, value := range arr {
// 		if counts[value] == nil {
// 			counts[value] = 0
// 		}
// 		counts[arr[i]]++
// 	}
// 	for value := range counts {
// 		if counts[value] > maxCount {
// 			maxCount = counts[value]
// 			maxNum = value
// 		}
// 	}
// 	return maxNum
// }

//1. Write a function which takes an array of numbers and returns the mode of that array.
// [2, 4, 4, 4, 6, 6, 1] - > 4
// [0, 0, 0, 10] -> 0
// [1] -> 1

// 2. Write a funciton which, given a year as a number, returns whether that number is a leap year.
// 2000 -> true
// 1900 -> false
// 2016 -> true
//writing a function
//checking if divisible by 4, but every 5th year we skip
//checking if divisible by 100 (as per James's astronomical observations)
//check if divisible by 400 (as per James's astronomical observations)
//if those are true, then return wthether the year is a leap year

func isLeapYear(year int) bool {
	if (year%4 == 0) && ((year%100 != 0) || (year%400 == 0)) {
		fmt.Printf("%d is a leap year! \n", year)
		return true
	}
	fmt.Printf("%d is not a leap year :( \n", year)
	return false
}

//3. Write a function which takes an array of integers > 0 and returns the first integer which does not appear in that array.
// [1,2,5] -> 3
// [1,2,3,4,5] -> 6
// [3,4,5] -> 1

func intNotInArr(arr []int) int {
	for i := range arr {
		if (i + 1) != arr[i] {
			return i + 1
		}
	}
	return len(arr) + 1
}
