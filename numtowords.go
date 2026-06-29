// Package numtowords contains methods to converts numbers to words.
// The first version converts positive integer number only, from 0
// to the MaxNum constant.
//
// No additional formatting options are provided
package numtowords

import "fmt"

// package level variables are allocated memory once the package is imported
// packages are loaded in natural order..

// MaxNum is the largest number that can be converted to words.
const MaxNum = 999

// Convert converts the specific number of type int to words.
func Convert(number int) (string, error) {
	if number > MaxNum || number < 0 {
		return "", fmt.Errorf("can only convert numbers between %v and %v", 0, MaxNum)
	}

	if number == 0 {
		return "zero", nil
	}

	units := [20]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}

	tens := [8]string{
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}
	result := ""

	if number > 99 {
		hundredindex := number / 100
		result += units[hundredindex] + " hundred"
		number = number % 100

		if number == 0 {
			return result, nil
		}

		result += " and "
	}

	if number > 19 {
		tensindex := number/10 - 2
		result += tens[tensindex]
		number %= 10

		if number == 0 {
			return result, nil
		}

		result += " "
	}

	result += units[number]

	return result, nil
}
