package main

import "testing"

funt TestEvenOrOdd(t *testing.T) {
	resulet := EvenOrOdd(10)
	if result != "even" {
		t.Errorf("expected: even, actual: %s", result)
	}