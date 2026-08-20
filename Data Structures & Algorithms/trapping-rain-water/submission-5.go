func trap(height []int) int {
	cur := 0
	res := 0
	var boundary int
	for {
		next, sub := findClosestGreater(height, cur)
		if next == -1 {
			break
		}
		res += max(min(height[next], height[cur]) * (next-cur-1) - sub, 0)
		cur = next
	}
	boundary = cur
	cur = len(height)-1
	fmt.Println(boundary, ", ", cur, ", ", res)
	fmt.Println("-----")
	for {
		next, sub := findClosestGreaterReverse(height, cur, boundary)
		if next == -1 {
			return res
		}
		res += max(min(height[next], height[cur]) * (cur-next-1) - sub, 0)
		cur = next
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

func findClosestGreaterReverse(height []int, cur int, minLeft int) (int, int) {
	next := -1
	sub := 0

	for i := cur-1; i >= minLeft; i-- {
		if height[i] >= height[cur] {
			return i, sub
		}
		sub += height[i]
	}
	return next, sub
}

// func findGreatestSmaller(height []int, cur int) (int, int) {
// 	next := -1
// 	sub, tmp := 0, 0
// 	minDiff := math.MaxInt
// 	for i := cur + 1; i < len(height); i++ {
// 		if height[i] < height[cur] {
// 			diff := height[cur] - height[i]
// 			if diff <= minDiff {
// 				minDiff = diff
// 				next = i
// 				sub = tmp
// 			}
// 		}
// 		tmp += height[i]
// 	}
// 	return next, sub
// }