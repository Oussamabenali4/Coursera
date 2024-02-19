package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var numList []int
	printSlice(numList)

	// append works on nil slices.
	for {
		var s string
		fmt.Println("Please enter distinct intergers:\n")
		fmt.Scan(&s)
		if s == "X" {
			break
		}

		//parseInt after input
		num, _ := strconv.Atoi(s)

		//If number is new, not existing in list, add
		if !IsContains(numList, num) {
			//Slice
			numList = append(numList, num)
			sort.Ints(numList)
			printSlice(numList)
		}
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func IsContains(nl []int, n int) bool {
	for _, i := range nl {
		if i == n {
			return true
		}
	}
	return false
}
