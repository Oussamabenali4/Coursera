package main

import (
	"fmt"
	"strconv"
)

func main() {
	sl := make([]int, 3)
	var s string
	var i int = 0
	for {
		fmt.Println("Add a number OR X to quit!")
		fmt.Scan(&s)
		if s == "x" || s == "X" {
			break
		} else {

			number, err := strconv.Atoi(s)
			if err == nil {
				if i < 3 {
					// append the first 3 places.
					sl[i] = number
					fmt.Println("The sorted slice is :")
					fmt.Println(sl)
					i++
				} else {
					sl = Sorting(append(sl, number))
					fmt.Println("The sorted slice is :")
					fmt.Println(sl)
				}

			} else {
				fmt.Println("Invalid input!")

			}

		}
	}

}

func Sorting(sl []int) []int {
	for i := 0; i < len(sl)-1; i++ {
		for j := i + 1; j < len(sl); j++ {
			if sl[i] > sl[j] {
				sl[i], sl[j] = sl[j], sl[i]
			}
		}
	}
	return sl
}
