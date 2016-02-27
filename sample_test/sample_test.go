package main

import "testing"

func TestSup(t *testing.T) {
	expected := "Sup, Ewan"
	outcome := sup("Ewan")
	if outcome != expected {
		t.Fatalf("Expected %s got %s", expected, outcome)
	}
}
