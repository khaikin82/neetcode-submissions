func removeElement(nums []int, val int) int {
   i := 0
   j := len(nums) - 1
   for i <= j {
	if nums[j] == val {
		j--
		continue
	}
	if nums[i] == val {
		nums[i] = nums[j]
		j--
	}
	i++
   }
   return j+1
}
