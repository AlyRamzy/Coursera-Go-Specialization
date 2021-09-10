package main

import "fmt"

func main() {
	fmt.Println("Please enter a floating point number : ")
	var input float64
	num, error := fmt.Scan(&input)

	if num == 0 || error != nil {
		fmt.Println("Invalid Input")
	} else {
		fmt.Printf("%d", int(input))
	}
}
