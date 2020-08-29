package constants

// TOTAL_KEYWORDS - total defined keywords
const TOTALKEYWORDS = 17

// https://www.programiz.com/c-programming/list-all-keywords-c-language
const AUTO = "auto"
const BREAK = "break"
const CASE = "case"
const CHAR = "char"
const CONST = "const"
const CONTINUE = "continue"
const DEFAULT = "default"
const DO = "do"
const DOUBLE = "double"
const ELSE = "else"
const ENUM = "enum"
const EXTERN = "extern"
const FOR = "for"
const GOTO = "goto"
const IF = "if"
const INT = "int"
const LONG = "long"

// KEYWORDS are array of keywords
var KEYWORDS = [TOTALKEYWORDS]string{
	AUTO, BREAK, CASE, CHAR, CONST, CONTINUE, DEFAULT, DO, DOUBLE,
	ELSE, ENUM, EXTERN, FOR, GOTO, IF, INT, LONG,
}

// IsKeyword will check if given string is keyword or not
func IsKeyword(token string) bool {
	for _, keyword := range KEYWORDS {
		if keyword == token {
			return true
		}
	}
	return false
}
