package reverseinteger

import (
	"testing"
)

func TestReverseSimple(t *testing.T) {
	input := 123
	expected := 321
	got := Reverse(input)

	if got != expected {
		t.Errorf("case 1 fail: got %d, expected %d", got, expected)
	}
}

func TestWithNegative(t *testing.T) {
	input := -321
	expected := -123
	got := Reverse(input)
	if got != expected {
		t.Errorf("case 2 fail: got %d, expected %d", got, expected)
	}
}

func TestWithZeros(t *testing.T) {
	input := 120
	expected := 21
	got := Reverse(input)
	if got != expected {
		t.Errorf("case 3 fail: got %d, expected %d", got, expected)
	}
}

func TestLimits(t *testing.T) {
	input := 1534236469
	expected := 0
	got := Reverse(input)
	if got != expected {
		t.Errorf("case 4 fail: got %d, expected %d", got, expected)
	}
}
