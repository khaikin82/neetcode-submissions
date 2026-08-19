func maxArea(heights []int) int {
	maxArea := 0
	n := len(heights)
	for i := 0; i < n-1; i++ {
		for j := n-1; j > i; j-- {
			if heights[j] >= heights[i] {
				maxArea = max(maxArea, heights[i] * (j-i))
				break
			}
			maxArea = max(maxArea, heights[j] * (j-i))
		}
	}
	return maxArea
}
