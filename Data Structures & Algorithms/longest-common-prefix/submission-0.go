func longestCommonPrefix(strs []string) string {
    idx := 0
	res := ""
	for {
		var commonChar byte
		for i, str := range(strs) {
			if idx >= len(str) {
				return res
			}
			if i == 0 {
				commonChar = str[idx]
			} else if str[idx] != commonChar {
				return res
			}
		}
		res += string(commonChar)
		idx++
	}
	return res
}
