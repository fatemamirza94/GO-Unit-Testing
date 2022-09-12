package calculator_test

import (
	"calculator"
	"testing"
)

//simple testing
/*func TestCalculateIsArmstrong(t *testing.T) {
	t.Fail()
} */

type TestCase struct {
	value    int
	expected bool
	actual   bool
}

//unit testing
/*func TestCalculateIsArmstrong(t *testing.T) {
	testCase := TestCase{

		value:    371,
		expected: true,
	}

	testCase.actual = calculator.CalculateIsArmstrong(testCase.value)

	if testCase.actual != testCase.expected {
		t.Fail()
	}
}

func TestNegativeCalculatorIsArmstrong(t *testing.T) {

	testCase := TestCase{

		value:    350,
		expected: false,
	}

	testCase.actual = calculator.CalculateIsArmstrong(testCase.value)

	if testCase.actual != testCase.expected {
		t.Fail()
	}
}
*/

//testing with subtests
func TestCalculateIsArmstrong(t *testing.T) {
	t.Run("should return true for 371 ", func(t *testing.T) {
		testCase := TestCase{
			value:    371,
			expected: true,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

	t.Run("should return true for 370 ", func(t *testing.T) {
		testCase := TestCase{
			value:    370,
			expected: true,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
}

func TestNegativeCalculatorIsArmstrong(t *testing.T) {

	testCase := TestCase{

		value:    350,
		expected: false,
	}

	testCase.actual = calculator.CalculateIsArmstrong(testCase.value)

	if testCase.actual != testCase.expected {
		t.Fail()
	}
}
