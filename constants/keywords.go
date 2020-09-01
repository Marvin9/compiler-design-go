package constants

// TOTAL_KEYWORDS - total defined keywords
const TOTALKEYWORDS = 32

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
const FLOAT = "float"
const FOR = "for"
const GOTO = "goto"
const IF = "if"
const INT = "int"
const LONG = "long"
const REGISTER = "register"
const RETURN = "return"
const SHORT = "short"
const SIGNED = "signed"
const SIZEOF = "sizeof"
const STATIC = "static"
const STRUCT = "struct"
const SWITCH = "switch"
const TYPEDEF = "typedef"
const UNION = "union"
const UNSIGNED = "unsigned"
const VOID = "void"
const VOLATILE = "volatile"
const WHILE = "while"

// KEYWORDS are array of keywords
var KEYWORDS = []string{
	AUTO, BREAK, CASE, CHAR, CONST, CONTINUE, DEFAULT, DO, DOUBLE,
	ELSE, ENUM, EXTERN, FLOAT, FOR, GOTO, IF, INT, LONG, REGISTER,
	RETURN, SHORT, SIGNED, SIZEOF, STATIC, STRUCT,
	SWITCH, TYPEDEF, UNION, UNSIGNED, VOID, VOLATILE, WHILE,
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
