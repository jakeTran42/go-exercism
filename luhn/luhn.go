package luhn

import (
	"strconv"
	"strings"
)

// Valid : Inputs a number as a string and returns a bool as to
// whether or not it is valid according to the Luhn formula
func Valid(numberStr string) bool {
	numberStr = strings.Replace(numberStr, " ", "", -1)

	if _, err := strconv.Atoi(numberStr); err != nil {
		return false
	}

	if len(numberStr) < 2 {
		return false
	}

	sum := 0

	for index := range numberStr {
		if index%2 != 0 {
			numberChar := numberStr[len(numberStr)-index-1]
			number, _ := strconv.Atoi(string(numberChar))
			number *= 2

			if number > 9 {
				number -= 9
			}

			sum += number
		} else {
			numberChar := numberStr[len(numberStr)-index-1]
			number, _ := strconv.Atoi(string(numberChar))
			sum += number
		}
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
