func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		need := target - nums[i]
		if j, exists := seen[need]; exists {
			return []int{j, i}
		}
		seen[nums[i]] = i
	}
	return nil
}
