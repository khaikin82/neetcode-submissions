func longestConsecutive(nums []int) int {
	longest := 0
	numSet := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	for k, _ := range numSet {
		if _, ok := numSet[k-1]; ok {
			continue
		}
		length := 1
		for {
			if _, ok := numSet[k+length]; ok {
				length++
			} else {
				break
			}			
		}
		longest = max(longest, length)
	}
	return longest
}