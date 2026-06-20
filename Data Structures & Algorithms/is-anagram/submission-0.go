func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    size := int('z') - int('a') + 1
    count := make([]int, size)
    for i := 0; i < len(s); i++ {
        count[int(s[i])-'a']++
        count[int(t[i])-'a']--
    }
    for i := 0; i < size; i++ {
        if count[i] != 0 {
            return false
        }
    }
    return true
}
