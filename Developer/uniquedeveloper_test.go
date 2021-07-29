package Developer

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestUniqueDevelopers(t *testing.T) {

	developers := []Developer{
		{Name: "A", Age: 10},
		{Name: "B", Age: 20},
		{Name: "C", Age: 30},
		{Name: "A", Age: 40},
		{Name: "D", Age: 50},
		{Name: "C", Age: 60},
	}
	expected := []string{"A", "B", "C", "D"}
	uniqueDevelopers := UniqueDevelopers(developers)
	sort.Strings(uniqueDevelopers)
	assert.Equal(t, expected, uniqueDevelopers)
}
func TestNegativeUniqueDevelopers(t *testing.T) {

	developers := []Developer{
		{Name: "A", Age: 10},
		{Name: "B", Age: 20},
		{Name: "C", Age: 30},
		{Name: "A", Age: 40},
		{Name: "D", Age: 50},
		{Name: "C", Age: 60},
	}
	expected := []string{"A", "B", "D"}
	uniqueDevelopers := UniqueDevelopers(developers)
	sort.Strings(uniqueDevelopers)
	assert.NotEqual(t, expected, uniqueDevelopers)
}
