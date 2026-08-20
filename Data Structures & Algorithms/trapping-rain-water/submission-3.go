func trap(height []int) int {
	cur := 0
	res := 0
	for {
		next, sub := findClosestGreater(height, cur)
		if next == -1 {
			break
		}
		res += max(min(height[next], height[cur]) * (next-cur-1) - sub, 0)
		fmt.Println(next, ", ", res)
		cur = next
	}
	for {
		next, sub := findGreatestSmaller(height, cur)
		if next == -1 {
			return res
		}
		res += max(min(height[next], height[cur]) * (next-cur-1) - sub, 0)
		cur = next
		fmt.Println(next, ", ", res)
	}
	return res
}

func findClosestGreater(height []int, cur int) (int, int) {
	next := -1
	sub := 0

	for i := cur + 1; i < len(height); i++ {
		if height[i] >= height[cur] {
			return i, sub
		}
		sub += height[i]
	}
	return next, sub
}

func findGreatestSmaller(height []int, cur int) (int, int) {
	next := -1
	sub, tmp := 0, 0
	minDiff := math.MaxInt
	for i := cur + 1; i < len(height); i++ {
		if height[i] < height[cur] {
			diff := height[cur] - height[i]
			if diff <= minDiff {
				minDiff = diff
				next = i
				sub = tmp
			}
		}
		tmp += height[i]
	}
	return next, sub
}