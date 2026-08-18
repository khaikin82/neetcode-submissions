func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		left := numbers[l]
		right := numbers[r]
		if left + right == target {
			return []int{l+1, r+1}
		}
		if left + right < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}
