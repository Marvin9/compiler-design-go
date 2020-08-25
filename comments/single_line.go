package comments

// SingleLineComment - detects single line comment
func SingleLineComment(input string) bool {
	for idx, char := range input {
		if char == '/' && input[idx+1] == '/' {
			return true
		}
	}
	return false
}
