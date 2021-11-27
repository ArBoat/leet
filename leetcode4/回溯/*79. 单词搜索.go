// 有used 全局 用匿名函数
func exist(board [][]byte, word string) bool {
    dir:=[][]int{{0,1},{0,-1},{1,0},{-1,0}}
    used:=make([][]bool, len(board)) 
    for i:= range used {
        // IM: len(board[0]
        used[i] = make([]bool, len(board[0]))
    }
    var check func(i,j,k int) bool
    check = func(i, j, k int) bool {
        // IM : 终止条件2个
        if board[i][j] != word[k] {
            return false
        }
        if k == len(word) - 1 {
            return true
        }
        // backTrack
        used[i][j] = true
        // IM: _,v
        for _, v:= range dir {
            if ni, nj:= i+v[0], j+v[1]; ni>=0 && ni < len(board) &&
            nj>=0 && nj<len(board[0]) && !used[ni][nj] {
                // 条件
                if check(ni, nj, k+1) {
                    return true
                }
            }
        }
        // 重置
        used[i][j] = false
        return false
    }

    for i, row := range board {
        for j:= range row {
            if check(i, j, 0) {
                return true
            }
        }
    }
    // IM: 非true
    return false
}

/*
O(mn*3^L)
O(mn)
回溯
*/