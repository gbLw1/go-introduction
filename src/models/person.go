package person

import (
	"fmt"
	"gointro/src/utils"
)

type Person struct {
	Name string
	Age  int
}

func PrintPerson(p Person) {
	fmt.Println()
	fmt.Println("User: {")
	fmt.Printf("  Name: %s\n", utils.Coalesce(p.Name, "Anonymous"))
	fmt.Printf("  Age: %d\n", p.Age)
	fmt.Println("}")
}
