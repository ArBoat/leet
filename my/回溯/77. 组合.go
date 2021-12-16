func combine(n int, k int) [][]int {
    res:=[][]int{}
    path:=[]int{}
    var backtrack func(path []int, cur int)
    backtrack = func(path []int, cur int){
        if len(path) == k {
            tmp:=make([]int, k)
            copy(tmp, path)
            res = append(res, tmp)
            return
        }
        // IM: len(path) + n - i + 1 >= k, i没在path中
        //for i:= cur; i<= n ; i++ {
        for i:= cur; i<= len(path)+n+1-k ; i++ {
            backtrack(append(path,i),i+1)
        }
    }
    backtrack(path, 1)
    return res
}

/*
n-i+1 >= k-len(path)
*/