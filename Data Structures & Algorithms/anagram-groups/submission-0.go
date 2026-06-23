import "slices"
import "cmp"

func groupAnagrams(strs []string) [][]string {
	sortArr := make([][]string, len(strs))
	for i, str := range(strs) {
		sortArr[i] = make([]string, 2)
		runeArr := []rune(str)
		slices.SortFunc(runeArr, func (a, b rune) int {
			return cmp.Compare(a, b)
		})
		sortArr[i][0] = string(runeArr)
		sortArr[i][1] = str
	}
	slices.SortFunc(sortArr, func(a, b []string) int {
		if len(a) < 2 || len(b) < 2 {
			fmt.Println("Err: ", a, b)
			return 0
		}
		return cmp.Compare(a[0], b[0])
	})
	res := make([][]string, 0)
	arr := []string{sortArr[0][1]}
	fmt.Println(sortArr)
	for i := 1; i < len(sortArr); i++ {
		if sortArr[i][0] == sortArr[i-1][0] {
			arr = append(arr, sortArr[i][1])
		} else {
			res = append(res, arr)
			arr = []string{sortArr[i][1]}
		}
	}
	res = append(res, arr)
	return res
	
}
