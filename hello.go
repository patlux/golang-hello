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

func getWhatToSay(name string) string {
	patrick := Person{name}
	return fmt.Sprintf("Hello '%s'.\n", patrick.name)
}
