package numtowords_test

import (
	"testing"

	"github.com/venkat0724/numtowords"
)

func TestInvalidNumber(t *testing.T) {
	result, err := numtowords.Convert(numtowords.MaxNum + 1)

	if err == nil {
		t.Logf("expected error but got %v", result)
		t.Fail()
	}

	_, err = numtowords.Convert(numtowords.MinNum - 1)

	if err == nil {
		t.Log("expected error")
		t.Fail()
	}
}

func TestZero(t *testing.T) {
	result, err := numtowords.Convert(0)

	if err != nil {
		t.Log("Covert 0 to words returned an error")
	}

	if result != "zero" {
		t.Logf("expected 'zero', received '%v'", result)
	}
}

func TestUnits(t *testing.T) {
	result, err := numtowords.Convert(99)
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
	for index, value := range units {
		result, err = numtowords.Convert(index)

		if err != nil {
			t.Fail()
		}

		if result != value {
			t.Fail()
		}
	}
}

func TestTens(t *testing.T) {
	testcases := map[int]string{
		34: "thirty four",
		60: "sixty",
	}

	for k, v := range testcases {
		result, err := numtowords.Convert(k)
		if err != nil {
			t.Logf("error while conversion")
			t.Fail()
		}

		if result != v {
			t.Logf("result not matched")
			t.Fail()
		}

	}
}

func TestHundreds(t *testing.T) {
	testcases := map[int]string{
		109: "one hundred and nine",
		777: "seven hundred and seventy seven",
	}
	for k, v := range testcases {
		result, err := numtowords.Convert(k)
		if err != nil {
			t.Logf("error while conversion")
			t.Fail()
		}
		t.Log(result)
		if result != v {
			t.Logf("result not matched")
			t.Fail()
		}

	}
}

func TestNegativeNumbers(t *testing.T) {
	testcases := map[int]string{
		-1:   "minus one",
		-99:  "minus ninety nine",
		-200: "minus two hundred",
		-999: "minus nine hundred and ninety nine",
		-0:   "zero",
	}

	for k, v := range testcases {
		result, err := numtowords.Convert(k)
		if err != nil {
			t.Logf("error while conversion %v", err)
			t.Fail()
		}
		t.Logf("result is %v", result)
		if result != v {
			t.Logf("result not matched")
			t.Fail()
		}
	}
}
