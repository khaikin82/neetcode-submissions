import "slices"
func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int, 0)
	for _, num := range nums {
		countMap[num]++
	}
	arr := make([][2]int, 0)
	for num, count := range countMap {
		arr = append(arr, [2]int{num, count})
	}
	slices.SortFunc(arr, func(a, b [2]int) int {
		return -(a[1]-b[1])
	})
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, arr[i][0])
	}
	return res
}
