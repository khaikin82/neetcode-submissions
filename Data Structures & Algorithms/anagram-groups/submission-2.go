func groupAnagrams(strs []string) [][]string {
	myMap := make(map[[26]int][]string)
	for _, str := range(strs) {
		count := [26]int{}
		for _, c := range(str) {
			count[c-'a']++
		}
		myMap[count] = append(myMap[count], str)
	}

	res := make([][]string, 0)
	for _, arr := range(myMap) {
		res = append(res, arr)
	}
	return res
}
