package main

import "fmt"

func main() {

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(slice); i++ {
		var message string
		sliceElement := slice[i]
		if (sliceElement % 2) == 0 {
			message = "EVEN"
		} else {
			message = "ODD"
		}
		fmt.Printf("Value %d is %s \n", sliceElement, message)
	}

}
