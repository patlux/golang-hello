package main

import "testing"

func TestGetWhatToSay(t *testing.T) {
  given := getWhatToSay("Patrick")
  expected := "Hello 'Patrick'.\n"
  if given != expected {
    t.Errorf("got %q want %q", given, expected)
  }
}
