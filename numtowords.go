// Package numtowords contains methods to converts numbers to words.
// The first version converts positive or negative integer, from MinNum to MaxNum, where MinNum and MaxNum are constants defined in the package. The package can be extended to convert larger numbers, but for now it is limited.
//
// No additional formatting options are provided
package numtowords

import "fmt"

// package level variables are allocated memory once the package is imported
// packages are loaded in natural order..

// MaxNum is the largest number that can be converted to words.
const MaxNum = 999

// MinNum is the smallest number that can be converted to words.
const MinNum = -999

var units = [20]string{
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

var tens = [8]string{
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

func commonFunction(number int) (string, error) {
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

// Convert converts the specific number of type int to words.
func Convert(number int) (string, error) {
	if number > MaxNum || number < MinNum {
		return "", fmt.Errorf("can only convert numbers between %v and %v", MinNum, MaxNum)
	}

	if number == 0 || number == -0 {
		return "zero", nil
	}

	if number < 0 {
		result, err := commonFunction(-number)
		if err != nil {
			return "", err
		}
		return "minus " + result, nil
	}

	return commonFunction(number)
}
