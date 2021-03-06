package main

import "testing"

func TestGetWhatToSay(t *testing.T) {

	assertString := func(t testing.TB, given string, expected string) {
		t.Helper()
		if given != expected {
			t.Errorf("got %q want %q", given, expected)
		}
	}

	t.Run("Should say hello world if no string is supplied", func(t *testing.T) {
		given := getWhatToSay("", "English")
		expected := "Hello, 'World'.\n"
		assertString(t, given, expected)
	})

	t.Run("Should say hello to people", func(t *testing.T) {
		given := getWhatToSay("Patrick", "English")
		expected := "Hello, 'Patrick'.\n"
		assertString(t, given, expected)
	})

	t.Run("Should say hello to people in german", func(t *testing.T) {
		given := getWhatToSay("Patrick", "German")
		expected := "Hallo, 'Patrick'.\n"
		assertString(t, given, expected)
	})

	t.Run("Should say hello to people in french", func(t *testing.T) {
		given := getWhatToSay("Patrick", "French")
		expected := "Bonjour, 'Patrick'.\n"
		assertString(t, given, expected)
	})

}
