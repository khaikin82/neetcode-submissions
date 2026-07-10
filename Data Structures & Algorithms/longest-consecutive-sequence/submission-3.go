import "slices"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slices.Sort(nums)

	longest := 1
	current := 1

	for i := 1; i < len(nums); i++ {
		switch {
		case nums[i] == nums[i-1]:
			continue

		case nums[i] == nums[i-1]+1:
			current++
			longest = max(longest, current)

		default:
			current = 1
		}
	}

	return longest
}