package dao

import (
	"testing"
)

//This is an integration test and could fail based on database state or connection
func TestGettingHighestCustomerId(t *testing.T) {
	total, err := GetHighestCustID()
	if err != nil {
		t.Error(err.Error())
	}
	expected := 1
	if total < expected {
		t.Errorf("highest customer id was incorrect, got: %d, want as least: %d.", total, expected)
	}
}

func TestRetrieveCustomers(t *testing.T) {
	total, err := RetrieveCustomers(0, 2)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(total)
	expectedSize := 1
	if len(total) < expectedSize {
		t.Errorf("size was incorrect, got: %d, want as least: %d.", total, expectedSize)
	}
}

func TestCreateAndUpdate(t *testing.T) {
	expectedName := "Carl"
	UpdateCustomer(Customer{Id: -1, First: expectedName})
	cust, _ := RetrieveCustomers(-1, 1)
	if cust[0].First != expectedName {
		t.Errorf("name was %s expected %s", cust[0].First, expectedName)
	}
}
