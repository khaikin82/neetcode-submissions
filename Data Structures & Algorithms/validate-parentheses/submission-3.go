func isValid(s string) bool {
    bracketMap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := bracketMap[c]; ok {
			stack = append(stack, c)
		} else {
			length := len(stack)
			if length == 0 || c != bracketMap[stack[length-1]] {
				return false
			}
			stack = stack[:length-1]
		}
	}
	return len(stack) == 0
}
