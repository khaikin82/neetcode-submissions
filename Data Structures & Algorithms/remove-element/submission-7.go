func removeElement(nums []int, val int) int {
    j := len(nums)-1
	
	for i := 0; i <= j; i++ {
		for j >= 0 && nums[j] == val {
			j--
		}
		if nums[i] == val && i <= j {
			nums[i] = nums[j]
			j--
		}
	}
	j++
	return j
}
