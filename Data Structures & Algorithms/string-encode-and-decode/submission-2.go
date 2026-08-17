type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var res strings.Builder
	for _, str := range(strs) {
		res.WriteString(strconv.Itoa(len(str)))
		res.WriteString("#")
		res.WriteString(str)
	}
	return res.String()
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}
	res := []string{}
	i := 0
	for i < len(encoded) {
		j := i
		for j < len(encoded) && encoded[j] != '#' {
			j++
		}
		size, _ := strconv.Atoi(encoded[i:j])
		i = j+1
		res = append(res, encoded[i:i+size])
		i = i+size
	}
	return res
}
