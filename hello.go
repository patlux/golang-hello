package main

import (
	"fmt"
)

type Person struct {
	name string
}

func main() {
	fmt.Printf(getWhatToSay("Patrick", "English"))
}

func getWhatToSay(name string, lang string) string {
	var person Person
	if name != "" {
		person = Person{name}
	} else {
		person = Person{name: "World"}
	}
	return fmt.Sprintf("%s '%s'.\n", getGreeting(lang), person.name)
}

func getGreeting(lang string) (prefix string) {

	switch lang {
	case "German":
		prefix = "Hallo,"
	case "French":
		prefix = "Bonjour,"
	default:
		prefix = "Hello,"
	}

	return
}
