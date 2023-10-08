package utils

import "fmt"

func CheckAge(age int) {
	fmt.Println()

	if age > 0 && age < 18 {
		fmt.Println("You are not old enough.")
	} else if age >= 18 {
		fmt.Println("You are old enough.")
	} else {
		fmt.Println("Invalid age.")
	}
}
