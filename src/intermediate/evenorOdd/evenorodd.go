package main

import (
	"errors"
	"fmt"
)

func main() {

	var (
		zero, even, odd, total int
		theError               error
	)

	var numArrry []int = []int{1, 2, 3, 4, -25, 5, 6, 7, 8, 9, 10, 0}

	/*for _, element := range numArrry {
		total++
		if element%2 == 0 {
			even++

		} else {
			odd++
		}

	} */

Abort:
	for _, element := range numArrry {

		switch {
		case element%2 == 0:
			even++
		case element == 0:
			zero++
		case element%2 == 1:
			odd++
		default:
			theError = errors.New(" Got negative number, SO Quitting")
			break Abort
		}
	}
	if theError != nil {
		fmt.Println(theError)
	}

	fmt.Printf(" Even : %d odd : %d total : %d", even, odd, total)
}
