func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}
	l := 0
	for ; l < n-1; l++ {
		if arr[l] >= arr[l+1] {
			break
		}
	}
	r := n - 1
	for ; r > 0; r-- {
		if arr[r] >= arr[r-1] {
			break
		}
	}

	return l == r && l != 0 && l != n-1
}