package fizzbuzz

import (
	"strconv"
)

// ItoFizzBuzz convert integer to fizzbuzz
func ItoFizzBuzz(number int) string {
	switch {
	case number%15 == 0:
		return "FizzBuzz"
	case number%3 == 0:
		return "Fizz"
	case number%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(number)
	}
}
