func lengthOfLongestSubstring(s string) int {
    win:=make(map[rune]int)
    left:=0
    rs:=[]rune(s)
    res:=0
    for right, v:= range rs {
        win[v]++
        for ;win[v]>1;left++{
            win[rs[left]]--
        }
        res = max(res, right-left+1)
    }
    return res
}

func max(a, b int) int {
    if a<b {
        return b
    }
    return a
}