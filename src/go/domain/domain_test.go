package domain

import "testing"

func TestHackyAmountHappyCase(t *testing.T) {
	units := uint(10)
	currency := Pound

	actual, err := NewHackyAmount(units, currency)
	if err != nil {
		t.Fatalf("failed to create account: %v", err)
	}

	expected := &HackyAmount{
		NumberOfSmallestWholeUnit: units,
		Currency:                  currency,
	}

	if *actual != *expected {
		t.Fatalf(`incorrect account. Expected: %v but received %v`, expected, actual)
	}
}
