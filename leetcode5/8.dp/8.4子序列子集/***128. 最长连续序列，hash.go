func longestConsecutive(nums []int) int {
	hash := make(map[int]bool)
	res := 0
	for _, v := range nums {
		hash[v] = true
	}
	for i := range hash {
		if !hash[i-1] {
			cur := i
			curLen := 1
			for hash[cur+1] { //IM：another index
				cur++
				curLen++
			}
			if curLen > res {
				res = curLen
			}
		}
	}
	return res
}

/*
hash map[int]bool
O(n)/O(n)
*/