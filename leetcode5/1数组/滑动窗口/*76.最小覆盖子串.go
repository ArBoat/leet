func minWindow(s string, t string) string {
	win, hash := make(map[byte]int), make(map[byte]int)
	left, right := 0, 0
	match, size := 0, len(s)+1 // 判空用
	for i := range t {
		hash[t[i]]++
	}
	i := 0             // 指针1
	for j := range s { // 指针2
		v := s[j]
		if _, ok := hash[v]; ok {
			win[v]++
			if win[v] == hash[v] {
				match++
			}
		}
		// len(hash)
		for ; match == len(hash); i++ {
			if j-i+1 < size {
				size = j - i + 1
				left, right = i, j
			}
			v := s[i]
			if _, ok := hash[v]; ok {
				win[v]--
				if win[v] < hash[v] {
					match--
				}
			}
		}
	}
	if size == len(s)+1 {
		return ""
	}
	return s[left : right+1]

}