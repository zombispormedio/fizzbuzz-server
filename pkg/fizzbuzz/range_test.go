package fizzbuzz

import (
	"testing"
)

func TestCreateFizzbuzzRange(t *testing.T) {
	values := CreateFizzbuzzRange(1, 5)

	if values[2] != "Fizz" {
		t.Error("Third element is not fizz")
	}

	if values[4] != "Buzz" {
		t.Error("Fifth element is not buzz")
	}
}
