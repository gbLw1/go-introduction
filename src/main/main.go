package main

import (
	"fmt"
	m "gointro/src/models" // alias for 'models'
	"gointro/src/utils"
	// "strconv"
)

func main() {
	fmt.Println("Hello, world!")

	name := askName()
	age := askAge()

	utils.CheckAge(age)

	user := m.Person{
		Name: name,
		Age:  age,
	}

	m.PrintPerson(user)

	// loopTest()
}

func askName() string {
	var name string

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	return name
}

func askAge() int {
	var age int

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	// check if age is valid
	if age < 0 {
		fmt.Println("Invalid age. Please enter a positive value.")
		return 0
	}

	return age
}

func loopTest() {
	fmt.Println()
	fmt.Println("Looping... from 1 to 10: ")

	var loopResult string

	for i := 0; i < 10; i++ {
		if i > 0 {
			loopResult += ", "
		}

		loopResult += fmt.Sprintf("%d", i+1)
		// loopResult += strconv.Itoa(i + 1) // lib 'strconv' also converts int to string
	}

	fmt.Println(loopResult)
}
