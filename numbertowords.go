package numbertowords

import (
	"errors"
)

//units
var words = [21]string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fiveteen",
	"sixteen", "seventee", "eightee", "nineteen",
}

//tens
var tens = [10]string{
	"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

//maxNumber is the largest number that this version
// of numberstowords can convert into words.
const maxNumber = 100

//Convert converts an integer between 0 and MaxNumber to words.
//If passed a number outside the valid range, it returns an error.
func Convert(number int) (string, error) {

	if number < 0 || number > maxNumber {
		return "", errors.New("number not in valid range")
	}

	if number == 100 {
		return "Hundred", nil
	}

	if number == 0 {
		return words[0], nil
	}

	var result string

	if number >= 20 && number < 100 {
		var r = number / 10
		result = result + tens[r]
		number = number % 10
	}

	if number >= 1 && number <= 19 {
		result = result + " " + words[number]
	}

	return result, nil

}
