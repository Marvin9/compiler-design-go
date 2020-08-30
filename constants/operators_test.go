package constants_test

import (
	"testing"

	"github.com/Marvin9/compiler-design-go/constants"
)

func TestOperatorsArray(t *testing.T) {
	operatorsArrayLen := len(constants.OPERATORS)
	totalOperators := constants.TOTALOPERATORS
	if operatorsArrayLen != totalOperators {
		t.Errorf("Total defined operators were %v, but In array found only %v.", totalOperators, operatorsArrayLen)
	}
}

func TestIsOperator(t *testing.T) {
	for _, operator := range constants.OPERATORS {
		if !constants.IsOperator(operator) {
			t.Errorf("%v operator is in array, but IsOperator returned false.", operator)
		}
	}

	randoms := []string{"foo", "bar", "@", "_", "blah", ""}
	for _, random := range randoms {
		if constants.IsOperator(random) {
			t.Errorf("%v is not operator, but IsOperator returned true", random)
		}
	}
}
