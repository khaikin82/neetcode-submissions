func longestCommonPrefix(strs []string) string {
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		res = compare(res, strs[i])
		fmt.Println(res)
	}
	return res
}

func compare(one, two string) string {
	var less int
	var res string
	if len(one) < len(two) {
		less = len(one)
		res = one
	} else {
		less = len(two)
		res = two
	}
	
	for i := 0; i < less; i++ {
		if one[i] != two[i] {
			return one[:i]
		}
	}
	return res
}
