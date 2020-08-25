package comments_test

import (
	"compiler-design-go/comments"
	"fmt"
	"testing"
)

func commentErr(input string, expected, got bool, t *testing.T) {
	var phrase string
	if expected {
		phrase = "contains"
	} else {
		phrase = "do not contain"
	}
	t.Errorf(fmt.Sprintf("%v %v comment(s) but got opposite", input, phrase))
}

type testCase struct {
	input           string
	expectedComment bool
}

func checkTestCases(testCases []testCase, checker func(string) bool, t *testing.T) {
	for _, tc := range testCases {
		if isComment := checker(tc.input); isComment != tc.expectedComment {
			commentErr(tc.input, tc.expectedComment, isComment, t)
		}
	}
}

var testCasesSingleLineComment = []testCase{
	testCase{
		"foo bar", false,
	},
	testCase{
		"foo // bar", true,
	},
	testCase{
		"foo / bar", false,
	},
}

func TestSingleLineComments(t *testing.T) {
	checkTestCases(testCasesSingleLineComment, comments.SingleLineComment, t)
}

var testCasesMultiLineComment = []testCase{
	testCase{
		`foo bar`, false,
	},
	testCase{
		`foo /*
		bar */ b`, true,
	},
	testCase{
		`foo /*
		
		bar`, false,
	},
}

func TestMultiLineComments(t *testing.T) {
	checkTestCases(testCasesMultiLineComment, comments.MultiLineComments, t)
}
