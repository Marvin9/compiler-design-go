package constants

const TOTALOPERATORS = 5

const PLUS = "+"
const MINUS = "-"
const MULTIPLICATION = "*"
const DIVISION = "/"
const MODULO = "%"

// OPERATORS - list of valid operators
var OPERATORS = [TOTALOPERATORS]string{
	PLUS, MINUS, MULTIPLICATION, DIVISION, MODULO,
}

// IsOperator will check if token is operator or not
func IsOperator(token string) bool {
	for _, operator := range OPERATORS {
		if operator == token {
			return true
		}
	}
	return false
}
