import "slices"
import "cmp"

func groupAnagrams(strs []string) [][]string {
	myMap := make(map[string][]string)
	for _, str := range(strs) {
		sortStr := sortString(str)
		myMap[sortStr] = append(myMap[sortStr], str)
	}

	res := make([][]string, 0)
	for _, arr := range(myMap) {
		res = append(res, arr)
	}
	return res
}

func sortString(str string) string {
	arr := []rune(str)
	slices.SortFunc(arr, func(a, b rune) int {
		return cmp.Compare(a, b)
	})
	return string(arr)
}
