func commonChars(words []string) []string {
	min := [26]int{}
	for i := range min {
		min[i] = math.MaxInt64
	}
	for _, word := range words {
		freq := [26]int{}
		for _, v := range word {
			freq[v-'a']++
		}
		for i, v := range freq {
			if v < min[i] {
				min[i] = v
			}
		}
	}
	res := []string{}
	for i, v := range min {
		for j := 0; j < v; j++ {
			res = append(res, string(i+'a'))
		}
	}
	return res
}