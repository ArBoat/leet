func letterCombinations(digits string) []string {
    res:=[]string{}
    // important
    if digits == "" { return res }
    hash:=make(map[string][]string)
    hash["2"]=[]string{"a","b","c"}
    hash["3"]=[]string{"d","e","f"}
    hash["4"]=[]string{"g","h","i"}
    hash["5"]=[]string{"j","k","l"}
    hash["6"]=[]string{"m","n","o"}
    hash["7"]=[]string{"p","q","r","s"}
    hash["8"]=[]string{"t","u","v"}
    hash["9"]=[]string{"w","x","y","z"}

    var backtrack func(path string, cur int)
    backtrack = func(path string, cur int){
        // 终止条件 IM
        // if cur == len(digits) {
        //     res = append(res, path)
        //     return 
        // }
        if len(path) == len(digits) {
            res = append(res, path)
            return
        }
        letters:=hash[string(digits[cur])]
        for _, v:= range letters {
            // 选择， 回退
            backtrack(path+v,cur+1)
        }
    }

    backtrack("",0)
    return res
}

/*
* m 是输入中对应 3 个字母的数字个数，n 是输入中对应 4 个字母的数字个数
* 时间复杂度：O((3^m+4^n)
* 空间复杂度：O(m+n)
backtrack
注意hash类型 和sting遍历类型
以及 边界
*/