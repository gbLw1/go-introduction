package main

import (
	"fmt"
	p "gointro/src/person"
	// "strconv"
)

func main() {
	fmt.Println("Hello, world!")

	name := askName()
	age := askAge()

	checkAge(age)

	user := p.Person{
		Name: name,
		Age:  age,
	}

	printUsuario(user)

	loopTest()
}

func coalesce(str, defaultValue string) string {
	if str == "" {
		return defaultValue
	}

	return str
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

func checkAge(age int) {
	fmt.Println()

	if age > 0 && age < 18 {
		fmt.Println("You are not old enough.")
	} else if age >= 18 {
		fmt.Println("You are old enough.")
	}
}

func printUsuario(user p.Person) {
	fmt.Println()
	fmt.Println("User: {")
	fmt.Printf("  Name: %s\n", coalesce(user.Name, "Anonymous"))
	fmt.Printf("  Age: %d\n", user.Age)
	fmt.Println("}")
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
