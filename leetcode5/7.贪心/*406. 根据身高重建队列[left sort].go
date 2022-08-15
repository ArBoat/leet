func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	res := make([][]int, 0)
	for _, person := range people {
		k := person[1]
		res = append(res, person)
		copy(res[k+1:], res[k:])
		res[k] = person
		// k:=person[1]
		// res = append(res[:k], append([][]int{person}, res[k:]...)...)
	}
	return res
}

/*
h 从大到小排
k 从小到大排
*/