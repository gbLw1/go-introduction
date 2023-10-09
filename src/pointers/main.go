package main

import "fmt"

func main() {
	rngNumber := 420
	fmt.Println("Print main:", rngNumber, &rngNumber)

	PrintReference(&rngNumber)
	PrintReference(nil)

	PrintAsCopy(rngNumber)

	ptr := &rngNumber // create a pointer to rngNumber

	*ptr = 500 // change the value of the pointer

	fmt.Println("Print main changed:", rngNumber, &rngNumber)
}

func PrintReference(value *int) {
	if value == nil {
		fmt.Println("No pointer provided")
		return
	}

	*value = 1337
	fmt.Println("Print reference changed:", *value, value)
}

func PrintAsCopy(value int) {
	fmt.Println("Print as copy:", value, &value)
	value = 69
	fmt.Println("Print as copy changed:", value, &value)
}

/*
	& - memory address (pointer)
	* - dereference (value of a pointer)
*/
