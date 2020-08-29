package constants_test

import (
	"testing"

	"github.com/Marvin9/compiler-design-go/constants"
)

func TestKeywordsArray(t *testing.T) {
	keywordsArrayLen := len(constants.KEYWORDS)
	totalKeywords := constants.TOTALKEYWORDS
	if keywordsArrayLen != totalKeywords {
		t.Errorf("Total defined keywords were %v, but In array found only %v.", totalKeywords, keywordsArrayLen)
	}
}

func TestIsKeyword(t *testing.T) {
	for _, keyword := range constants.KEYWORDS {
		if !constants.IsKeyword(keyword) {
			t.Errorf("%v keyword is in array, but IsKeyword returned false.", keyword)
		}
	}

	randoms := []string{"foo", "bar", "@", "_", "blah", ""}
	for _, random := range randoms {
		if constants.IsKeyword(random) {
			t.Errorf("%v is not keyword, but IsKeyword returned true", random)
		}
	}
}
