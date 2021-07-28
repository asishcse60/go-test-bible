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
	testCase := TestCase{
		value:    371,
		expected: true,
	}

	testCase.actual = CalculateIsArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}
func TestNegativeCalculateIsArmstrong(t *testing.T) {
	testCase := TestCase{
		value: 372,
		expected: false,
	}

	testCase.actual = CalculateIsArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}