func longestConsecutive(nums []int) int {
	hset := make(map[int]struct{}, 0)
	for _, num := range(nums) {
		hset[num] = struct{}{}
	}
	res := 0
	for _, num := range(nums) {
		if _, ok := hset[num-1]; ok {
			continue
		}
		count := 1
		for {
			if _, exists := hset[num+1]; exists {
				count++
				num++
			} else {
				res = max(res, count)
				break
			}
		}
	}
	return res
}
