func isPalindrome(s string) bool {
	i := 0
	j := len(s)-1
	s = strings.ToUpper(s)
	for i < j {
		for i < j && !isUpperLetterOrNumber(s[i]) {
			i++
		}
		for i < j && !isUpperLetterOrNumber(s[j]) {
			j--
		}
		if i >= j {
			return true
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isUpperLetterOrNumber(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
