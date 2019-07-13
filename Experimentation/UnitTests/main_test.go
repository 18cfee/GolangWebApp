package main

import "testing"

func TestAddition(t *testing.T) {
	total := Sum(5, 5)
	expected := 11
	if total != expected {
		t.Errorf("sum was incorrect, got: %d, want: %d.", total, expected)
	}
}
