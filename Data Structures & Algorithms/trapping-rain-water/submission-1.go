func trap(height []int) int {
	cur := 0
	next := 0
	res := 0
	for {
		next = findClosestGreater(height, cur)
		if next != -1 {
			res += totalAmountBetween(height, cur, next)
			fmt.Println(next, ", ", res)
			cur = next
		} else {
			for {
				next = findGreatestSmaller(height, cur)
				if next == -1 {
					return res
				}
				res += totalAmountBetween(height, cur, next)
				cur = next
				fmt.Println(next, ", ", res)
			}
		}
	}
	return res
}

func findClosestGreater(height []int, cur int) int {
	next := -1

	for i := cur + 1; i < len(height); i++ {
		if height[i] >= height[cur] {
			return i
		}
	}
	return next
}

func findGreatestSmaller(height []int, cur int) int {
	next := -1
	minDiff := math.MaxInt
	for i := cur + 1; i < len(height); i++ {
		if height[i] < height[cur] {
			diff := height[cur] - height[i]
			if diff <= minDiff {
				minDiff = diff
				next = i
			}
		}
	}
	return next
}

func totalAmountBetween(height []int, cur, next int) int {
	sub := 0
	for i := cur+1; i < next; i++ {
		sub += height[i]
	}
	return max(min(height[next], height[cur]) * (next-cur-1) - sub, 0)
}