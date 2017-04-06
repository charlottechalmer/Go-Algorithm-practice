package main

import (
	"fmt"
	"sort"
)

func main() {
	test1 := []int{1, 3, 4, 4, 4, 5, 3, 5}
	fmt.Println(mode(test1))

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

	test8 := []int{1, 1, 2, 3, 1, 2, 3}
	fmt.Println(arrayWithoutRepeats(test8))

	test9 := []int{2, 3, 4}
	test95 := []int{1, 2, 3}
	fmt.Println(isSameValues(test9, test95))
	test10 := []int{1, 1, 1}
	test105 := []int{1, 1, 1, 1}
	fmt.Println(isSameValues(test10, test105))
	test11 := []int{3, 2, 4}
	fmt.Println(isSameValues(test9, test11))

	test12 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arrayAdjustedWithCutoff(test12, 4))
	// --> [1,2,3,4,4,4,4,4]
	test13 := []int{1, 5, 7, 3, 1, 5, 7}
	fmt.Println(arrayAdjustedWithCutoff(test13, 3))
	// --> [1,1,3,3,3,3,3]

	test14 := []int{1, 2, 5, 6, 9}
	test145 := []int{3, 4, 5, 10}
	fmt.Println(combinedSortedArr(test14, test145))
	//--> [1,2,3,4,5,5,6,9,10]
	test15 := []int{-1, 0, 1}
	test155 := []int{-2, 2}
	fmt.Println(combinedSortedArr(test15, test155))
	//--> [-2, -1, 0, 1, 2]
}

//1. Write a function which takes an array of numbers and returns the mode of that array.
// [2, 4, 4, 4, 6, 6, 1] - > 4
// [0, 0, 0, 10] -> 0
// [1] -> 1

func mode(arr []int) int { //passing in an array of ints and returning an int(the most frequent occurence)
	//map values seen and increment when seen another
	occurenceMap := make(map[int]int)
	var ok bool
	var counter int
	var maxNumber = 1
	for _, arrVal := range arr {
		mapKey := arrVal
		counter := occurenceMap[mapKey]
		//looping through the array, looking at the values
		if arrVal, ok = occurenceMap[mapKey]; ok {
			//if the array value is in the map as a key, then increment the map value
			fmt.Printf("%v is in the map \n", mapKey)
			counter++
			occurenceMap[mapKey] = counter
			fmt.Printf("%v is counter \n", counter)
		} else {
			//else, the map value (indicating the occurence) is 1
			fmt.Printf("%v is NOT in the map \n", mapKey)
			counter = 1
			occurenceMap[mapKey] = counter
		}
		fmt.Printf("Occurence map: %v \n", occurenceMap)
	}
	for key, value := range occurenceMap {
		//loop through the map
		if value > counter {
			counter = value
			maxNumber = key
		}
	}
	fmt.Printf("the most frequent occurring integer is: %v, occuring %v times \n", maxNumber, counter)
	return maxNumber
}

// 2. Write a funciton which, given a year as a number, returns whether that number is a leap year.
// 2000 -> true
// 1900 -> false
// 2016 -> true
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

//4. Write a function which takes an array of integers and returns an array with any duplicate integers removed.
// [1,1,2,3,1,2,3] -> [1,2,3]
// [1,4,4,4,2,2,4,4,4] -> [1,4,2]
func arrayWithoutRepeats(arr []int) []int {
	intsSeen := map[int]bool{}
	for i := range arr {
		intsSeen[arr[i]] = true
	}
	result := []int{}
	for key := range intsSeen {
		result = append(result, key)
	}
	return result
}

//Write a function which takes in two arrays and determines if they contain the same number of the same values.
// [], [] -> true
// [2, 3, 4], [1, 2, 3] -> false
// [1, 1, 1], [1, 1, 1, 1] -> false

func isSameLength(arr1 []int, arr2 []int) bool {
	if len(arr1) == len(arr2) {
		return true
	}
	return false
}

func isSameValues(arr1 []int, arr2 []int) bool {
	if isSameLength(arr1, arr2) == true {
		fmt.Printf("arrays are the same length \n")
		sort.Ints(arr1)
		fmt.Printf("arr1 sorted: %v \n", arr1)
		sort.Ints(arr2)
		fmt.Printf("arr2 sorted: %v \n", arr2)
		for i := 0; i < len(arr1); i++ {
			if arr1[i] != arr2[i] {
				return false
			}
		}
		return true
	}
	return false
}

//Write a function which takes in an array of numbers and a max cutoff value, and returns a new array with elements limited by the cutoff value.
// [1,2,3,4,5,6,7,8], 4 -> [1,2,3,4,4,4,4,4]
// [1,5,7,3,1,5,7], 3 -> [1,3,3,3,1,3,3]

func arrayAdjustedWithCutoff(arr []int, cutoffVal int) []int {
	sort.Ints(arr)
	for i := range arr {
		if arr[i] == cutoffVal {
			for j := i; j < len(arr); j++ {
				arr[j] = cutoffVal
			}
		}
	}
	return arr
}

//Write a function which takes two sorted lists and merges them into a new sorted list.
// [1,2,5,6,9], [3,4,5,10] -> [1,2,3,4,5,5,6,10]
// [], [] -> []
// [-1, 0, 1], [-2, 2] -> [-2, -1, 0, 1, 2]
func combinedSortedArr(arr1 []int, arr2 []int) []int {
	result := []int{}
	result = append(arr1, arr2...)
	sort.Ints(result)
	return result
}
