package calculator
import (
	"testing"

)

type TestCase struct {
	name string
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsArmstrong(t *testing.T) {
	t.Run("test for all 3 digit armstrong numbers", func(t *testing.T) {
		tests := []TestCase{
			TestCase{name:"test for 153",value: 153, expected: true},
			TestCase{name:"test for 370",value: 370, expected: true},
			TestCase{name:"test for 371",value: 371, expected: true},
			TestCase{name:"test for 407",value: 407, expected: true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual := CalculateIsArmstrong(test.value)
				if test.expected != actual {
					t.Fail()
				}
			})
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

func BenchmarkCalculateIsArmstrong370(b *testing.B) {
	benchmarkCalculateIsArmstrong(370, b)
}
func BenchmarkCalculateIsArmstrong371(b *testing.B) {
	benchmarkCalculateIsArmstrong(371, b)
}
func BenchmarkCalculateIsArmstrong0(b *testing.B) {
	benchmarkCalculateIsArmstrong(0, b)
}

func benchmarkCalculateIsArmstrong(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateIsArmstrong(input)
	}
}