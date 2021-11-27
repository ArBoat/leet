func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 { 
        return false
    }
    i,j:=len(matrix)-1, 0
    for j<=len(matrix[0])-1 && i >= 0 {
        if matrix[i][j] == target {
            return true
        } else if matrix[i][j] > target {
            i--
        } else {
            j++
        }
    }
    // IM: return false
    return false
}

/*
O(m+n)/O(1)
二维数组左下角,妙啊
*/