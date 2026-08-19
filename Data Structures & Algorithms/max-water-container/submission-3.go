func maxArea(heights []int) int {
    res := 0
	l, r := 0, len(heights)-1
	for l < r {
		res = max(res, min(heights[l], heights[r]) * (r-l))
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return res
}