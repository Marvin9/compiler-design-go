package validate_test

import (
	"testing"

	"github.com/Marvin9/compiler-design-go/constants"
	"github.com/Marvin9/compiler-design-go/validate"
)

func TestIsValidIdentifier(t *testing.T) {
	// empty string
	if validate.IsValidIdentifier("") {
		t.Errorf("Empty string is not valid identifier")
	}

	// keywords
	for _, keyword := range constants.KEYWORDS {
		if validate.IsValidIdentifier(keyword) {
			t.Errorf("%v is keyword and not valid identifier.", keyword)
		}
	}

	// invalid first character
	invalidFirstChars := []string{
		"1abc", "@abc", " abc", "-abc",
	}
	for _, invalidFirstChar := range invalidFirstChars {
		if validate.IsValidIdentifier(invalidFirstChar) {
			t.Errorf("%v has invalid first character, but returned as valid identifier", invalidFirstChar)
		}
	}

	// invalid identifier
	invalidIdentifiers := []string{
		"ab@", "", "a_b_c#$()", "+`!", "%$#@#",
	}
	for _, invalidIdentifier := range invalidIdentifiers {
		if validate.IsValidIdentifier(invalidIdentifier) {
			t.Errorf("%v is invalid identifier, but returned as valid identifier", invalidIdentifier)
		}
	}

	// valid identifiers
	validIdentifiers := []string{
		"abc", "_abc", "ab_c", "ab_c_d", "_123", "ab123",
	}
	for _, validIdentifier := range validIdentifiers {
		if !validate.IsValidIdentifier(validIdentifier) {
			t.Errorf("%v is valid identifier, but returned as invalid identifier", validIdentifier)
		}
	}
}
