func longestSubstring(s string, k int) int {
    return find(0, len(s)-1, k, s)
}

func find(start, end, k int, s string) int {
    if end-start+1<k {
        return 0
    }
    freq:=make(map[byte]int)
    for i:= start; i<=end; i++ {
        freq[s[i]]++
    }
    for end-start+1 >=k && freq[s[start]] < k {
        start++
    }
    for end-start+1 >=k && freq[s[end]] < k {
        end--
    }
    if end-start+1<k {
        return 0
    }
    for i:=start;i<=end;i++ {
        if freq[s[i]]<k{
            return max(find(start,i-1,k,s),find(i+1, end, k,s))
        }
    }
    return end-start+1
}
func max(a, b int ) int {
    if a>b{
        return a
    }
    return b
}

/*
分治加速
*/