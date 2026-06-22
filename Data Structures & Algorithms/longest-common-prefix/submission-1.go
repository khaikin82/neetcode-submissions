func longestCommonPrefix(strs []string) string {
    idx := 0
	var res strings.Builder
	for {
		var commonChar byte
		for i, str := range(strs) {
			if idx >= len(str) {
				return res.String()
			}
			if i == 0 {
				commonChar = str[idx]
			} else if str[idx] != commonChar {
				return res.String()
			}
		}
		res.WriteByte(commonChar)
		idx++
	}
	return res.String()
}
