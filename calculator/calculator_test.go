package calculator
import (
	"testing"

)

type TestCase struct {
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsArmstrong(t *testing.T) {

	t.Run("should return true for 371", func(t *testing.T) {
		testCase := TestCase{
			value:    371,
			expected: true,
		}
		testCase.actual = CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
	t.Run("should return true for 370", func(t *testing.T) {
		testCase := TestCase{
			value:    370,
			expected: true,
		}
		testCase.actual = CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

}
func TestNegativeCalculateIsArmstrong(t *testing.T) {
	t.Run("should return false for 350", func(t *testing.T) {
		testCase := TestCase{
			value: 372,
			expected: false,
		}

		testCase.actual = CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
	t.Run("should return false for 300", func(t *testing.T) {
		testCase := TestCase{
			value: 300,
			expected: false,
		}

		testCase.actual = CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

}