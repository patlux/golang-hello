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

	prefix := "Hello," // default

	if lang == "German" {
		prefix = "Hallo,"
	} else if lang == "French" {
		prefix = "Bonjour,"
	}

	return fmt.Sprintf("%s '%s'.\n", prefix, person.name)
}
