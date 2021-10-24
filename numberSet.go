package main

import (
	"fmt"
	"strconv"
)

// Creates a Type called numberSet
// Which is a slice of int32

type numberSet []int32

func createNewNumberSet(min int32, max int32) numberSet {
	numbers := numberSet{}

	for i := min; i < max; i++ {
		numbers = append(numbers, i)
	}

	return numbers
}

func oddOrEven(n int) string {
	if n%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func (n numberSet) printEvenOrOdd() {
	for i := range n {
		fmt.Println(strconv.Itoa(i) + " is " + oddOrEven(i))
	}
}
