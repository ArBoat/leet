//先match再判断长度
func findAnagrams(s string, p string) []int {
	res := []int{}
	win, hash := make(map[byte]int), make(map[byte]int)
	match := 0
	for i := range p {
		hash[p[i]]++
	}
	i := 0
	for j := range s {
		v := s[j]
		if _, ok := hash[v]; ok {
			win[v]++
			if win[v] == hash[v] {
				match++
			}
		}
		for ; match == len(hash); i++ {
			if j-i+1 == len(p) {
				res = append(res, i)
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
	return res
}

// 先判断长度再match
func findAnagrams(s string, p string) []int {
	res := []int{}
	win, hash := make(map[byte]int), make(map[byte]int)
	match := 0
	for i := range p {
		hash[p[i]]++
	}
	i := 0
	for j := 0; j < len(s); j++ {
		v := s[j]
		if _, ok := hash[v]; ok {
			win[v]++
			if win[v] == hash[v] {
				match++
			}
		}
		for ; j-i+1 >= len(p); i++ {
			if match == len(hash) {
				res = append(res, i)
			}
			v := s[i]
			if _, ok := hash[v]; ok {
				if win[v] == hash[v] {
					match--
				}
				// IM -- 顺序
				win[v]--
			}
		}
	}
	return res
}