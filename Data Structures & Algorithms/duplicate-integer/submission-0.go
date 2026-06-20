func hasDuplicate(nums []int) bool {
    countMap := make(map[int]int)
    for _, num := range(nums) {
        if _, ok := countMap[num]; ok {
            return true
        } else {
            countMap[num] = 0
        }
    }
    return false
}
