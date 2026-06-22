func longestCommonPrefix(strs []string) string {
	for idx := 0; idx < len(strs[0]); idx++ {
		for _, str := range(strs) {
			if idx == len(str) || str[idx] != strs[0][idx] {
				return strs[0][:idx]
			}
		}
	}
	return strs[0]
}
