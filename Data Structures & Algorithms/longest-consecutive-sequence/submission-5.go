func longestConsecutive(nums []int) int {
	hset := make(map[int]struct{}, 0)
	for _, num := range(nums) {
		hset[num] = struct{}{}
	}
	res := 0
	for _, num := range(nums) {
		_, ok := hset[num-1]
		if ok {
			continue
		}
		count := 1
		for {
			_, exists := hset[num+1]
			if exists {
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
