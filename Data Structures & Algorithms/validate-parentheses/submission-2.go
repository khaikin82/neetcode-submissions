func isValid(s string) bool {
    bracketMap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		c := s[i]
		_, ok := bracketMap[c]; 
		if ok {
			stack = append(stack, c)
		} else {
			length := len(stack)
			if length == 0 {
				return false
			}
			if c == bracketMap[stack[length-1]] {
				stack = stack[:length-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
