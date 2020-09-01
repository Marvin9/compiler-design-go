package constants

const TOTALOPERATORS = 22

// ARITHMETIC
const PLUS = "+"
const MINUS = "-"
const MULTIPLICATION = "*"
const DIVISION = "/"
const MODULO = "%"

// RELATIONAL
const EQUAL = "=="
const NOT_EQUAL = "!="
const GREATER_THAN = ">"
const LESS_THAN = ">"
const GREATER_THAN_EQUAL = ">="
const LESS_THAN_EQUAL = "<="

// BITWISE
const AND = "&"
const OR = "|"
const XOR = "^"
const TWOS_COMPLEMENT = "~"
const RIGHT_SHIFT = ">>"
const LEFT_SHIFT = "<<"

// ASSIGNMENT
const ASSIGN = "="
const ADD = "+="
const SUBTRACT = "-="
const MULTIPLY = "*="
const DIVIDE = "/="

// OPERATORS - list of valid operators
var OPERATORS = [TOTALOPERATORS]string{
	PLUS, MINUS, MULTIPLICATION, DIVISION, MODULO,
	EQUAL, NOT_EQUAL, GREATER_THAN, LESS_THAN,
	GREATER_THAN_EQUAL, LESS_THAN_EQUAL, AND, OR, XOR,
	TWOS_COMPLEMENT, RIGHT_SHIFT, LEFT_SHIFT,
	ASSIGN, ADD, SUBTRACT, MULTIPLY, DIVIDE,
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
