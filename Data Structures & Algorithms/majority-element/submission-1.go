func majorityElement(nums []int) int {
    res, count := 0, 0
	for _, num := range(nums) {
		if count == 0 {
			res = num
		}
		if res == num {
			count++
		} else {
			count--
		}
	}
	return res
}
