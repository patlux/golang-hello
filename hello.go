package main

import (
	"fmt"
)

type Person struct {
	name string
}

func main() {
	fmt.Printf(getWhatToSay("Patrick"))
}

const helloPrefix = "Hello,"

func getWhatToSay(name string) string {
	var person Person;
	if name != "" {
		person = Person{name}
	} else {
		person = Person{name: "World"}
	}
	return fmt.Sprintf("%s '%s'.\n", helloPrefix, person.name)
}
