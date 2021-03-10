package numbertowords_test

import (
	"testing"

	"github.com/rajeshtatipaka/numbertowords"
)

func TestInvalidInput(t *testing.T) {
	result, err := numbertowords.Convert(1)
	if result == "" || err != nil {
		t.Fatal("failed test for input -1")
	}

	result, err = numbertowords.Convert(100)
	if result == "" || err != nil {
		t.Fatal("failed test for input 1000")
	}
}

func TestUnits(t *testing.T) {
	testcases := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "siz",
		7: "seven",
		8: "eight",
		9: "nine",
	}
	for number := range testcases {
		_, err := numbertowords.Convert(number)
		if err != nil {
			t.Fatal("Failed for number: ", number, err)
		}
	}
}
