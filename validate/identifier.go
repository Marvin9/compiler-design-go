package validate

import (
	"unicode"

	"github.com/Marvin9/compiler-design-go/constants"
)

// IsValidIdentifier is used to check if token is valid identifier or not
// rules:
// 1. It should not be keyword
// 2. First char should not be anything except _, a-z, A-Z
// 3. Next chars should not be anything except _, a-z, A-Z, 0-9
func IsValidIdentifier(token string) bool {
	tokenLength := len(token)
	if tokenLength == 0 {
		return false
	}

	checkKeyword := constants.IsKeyword(token)
	if checkKeyword {
		return false
	}

	firstChar := token[0]

	// first character should not be anything except _, a-z, A-Z
	if !(firstChar == '_' || unicode.IsLetter(rune(firstChar))) {
		return false
	}

	for i := 1; i < tokenLength; i++ {
		tokenChar := rune(token[i])
		// first character afterwards, It should not be anything except _, a-z, A-Z, 0-9
		if !(tokenChar == '_' || unicode.IsDigit(tokenChar) || unicode.IsLetter(tokenChar)) {
			return false
		}
	}

	return true
}
