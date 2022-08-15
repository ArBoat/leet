type point struct {
	name    string
	visited bool
}

func findItinerary(tickets [][]string) []string {
	if len(tickets) == 0 {
		return nil
	}
	hash := make(map[string][]point)
	//key : from ; []point: to
	for _, v := range tickets {
		hash[v[0]] = append(hash[v[0]], point{name: v[1]})
	}
	for _, v := range hash {
		sort.Slice(v, func(a, b int) bool {
			return v[a].name < v[b].name
		})
	}

	path := []string{"JFK"}

	var backtrack func(start string) bool
	backtrack = func(start string) bool {
		if len(path) == len(tickets)+1 {
			return true
		}

		for i := 0; i < len(hash[start]); i++ {
			if !hash[start][i].visited {
				hash[start][i].visited = true
				path = append(path, hash[start][i].name)
				if backtrack(hash[start][i].name) {
					return true
				}
				path = path[:len(path)-1]
				hash[start][i].visited = false
			}
		}
		return false
	}
	backtrack("JFK")
	return path
}