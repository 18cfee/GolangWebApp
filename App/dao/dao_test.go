package dao

import (
	"testing"
)

//This is an integration test and could fail based on database state or connection
func TestGettingHighestCustomerId(t *testing.T) {
	total, err := GetHighestCustId()
	if err != nil {
		t.Error(err.Error())
	}
	expected := 1
	if total < expected {
		t.Errorf("highest customer id was incorrect, got: %d, want as least: %d.", total, expected)
	}
}

func TestRetrieveCustomers(t *testing.T) {
	total, err := RetrieveCustomers()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(total)
	expectedSize := 1
	if len(total) < expectedSize {
		t.Errorf("size was incorrect, got: %d, want as least: %d.", total, expectedSize)
	}
}
