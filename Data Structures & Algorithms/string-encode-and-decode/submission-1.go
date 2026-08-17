type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	lens := make([]string, len(strs))
	for i, str := range(strs) {
		lens[i] = strconv.Itoa(len(str))
	}
	return strings.Join(lens, ",") + "#" + strings.Join(strs, "")
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}
	parts := strings.SplitN(encoded, "#", 2)
	lens := strings.Split(parts[0], ",")
	strs := parts[1]
	m := 0
	res := make([]string, len(lens))
	for i, len := range(lens) {
		num, _ := strconv.Atoi(len)
		n := m + num
		res[i] = strs[m:n]
		m = n
	}
	return res
}
