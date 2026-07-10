import "slices"
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slices.Sort(nums)
	count := 1
	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] == nums[i-1] + 1 {
			count++
		} else {
			res = max(res, count)
			count = 1
		}
	}
	res = max(res, count)
	return res
}
