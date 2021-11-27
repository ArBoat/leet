func kthSmallest(matrix [][]int, k int) int {
    n:=len(matrix)
    l, r := matrix[0][0], matrix[n-1][n-1]
    for l < r {
        mid := l + (r-l)/2
        if check(matrix, k, mid) {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return l
}
// 分清matrix值和索引
// bool
func check(matrix [][]int, k, mid int) bool {
    n:= len(matrix)
    i,j:= n-1, 0
    count:=0
    for i>=0 && j < n {
        // <=
        if matrix[i][j]<= mid {
            count += i + 1
            j++
        } else {
            i--
        }
    }
    // >= k
    return count >= k
}

/*
O(logn)/O(1)
二分查找+有序矩阵左下角
*/