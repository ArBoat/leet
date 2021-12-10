func groupAnagrams(strs []string) [][]string {
	hash := make(map[[26]int][]string)
	for _, str := range strs {
		count := [26]int{}
		for _, v := range str {
			count[v-'a']++
		}
		hash[count] = append(hash[count], str)
	}
	res := make([][]string, 0)
	for _, v := range hash {
		res = append(res, v)
	}
	return res
}