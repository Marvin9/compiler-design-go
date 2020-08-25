package comments

// MultiLineComments - detects multiline comment
func MultiLineComments(input string) bool {
	limit := len(input)
	idx := 0
	for idx < limit {
		char := input[idx]
		if char == '/' && input[idx+1] == '*' {
			idx += 2
			for idx < limit {
				char = input[idx]
				if char == '*' && input[idx+1] == '/' {
					return true
				}
				idx++
			}
		}
		idx++
	}
	return false
}
