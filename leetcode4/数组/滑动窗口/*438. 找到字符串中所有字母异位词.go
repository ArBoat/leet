func findAnagrams(s string, p string) []int {
    win, hash:=make(map[rune]int), make(map[rune]int)
    match:=0
    left:=0
    rs:=[]rune(s)
    res:=[]int{}
    for _, v:= range p{
        hash[v]++
    }
    for right, v:= range rs {
        if _, ok:= hash[v];ok{
            win[v]++
            if win[v]==hash[v]{
                match++
            }
        }
        for ;right-left+1>=len(p);left++{
            if match == len(hash) {
                res = append(res, left)
            } 
            temp:=rs[left]
            if _, ok := hash[temp]; ok {
                if  win[temp] == hash[temp] {
                    match--
                }
                win[temp]--
            }
        }
    }
    return res
}

/*
O(n)/O(n)
滑动窗口+hash+双指针
*/