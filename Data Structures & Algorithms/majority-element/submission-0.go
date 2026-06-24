func majorityElement(nums []int) int {
    myMap := make(map[int]int)
	halfLen := int(len(nums) / 2)
	for _, num := range(nums) {
		myMap[num]++
		if myMap[num] > halfLen {
			return num
		}
	}
	return 0
}
