type inum struct {
    index int
    val int 
}

func countSmaller(nums []int) []int {
    if len(nums)==0  {
        return nil
    }
    hash:=make(map[int]int)
    inums:= make([]inum, len(nums))
    for i, v:= range nums{
        inums[i] = inum{i, v}
    }
    mergeSort(inums, 0, len(nums)-1, hash)
    res:=make([]int,len(nums))
    for k,v:=range hash {
        res[k]=v
    }
    // fmt.Println(nums)
    return res
}

func mergeSort(inums []inum, l, r int, hash map[int]int) {
    if l >= r{ return }
        m:= l + (r-l)/2
        mergeSort(inums, l, m, hash)
        mergeSort(inums, m+1, r, hash)
        tmp:=[]inum{}
        i, j:= l, m+1
        for i<=m && j<=r{
            if inums[i].val<=inums[j].val{
                tmp = append(tmp, inums[i])
                hash[inums[i].index] += j-(m+1)
                i++
            } else {
                tmp = append(tmp, inums[j])
                j++
            }
        }
        for ;i<=m;i++ {
            tmp = append(tmp, inums[i])
            hash[inums[i].index] += r - (m+1) + 1
        }
        for ;j<=r;j++{
            tmp = append(tmp, inums[j])
        }
        for k:=l;k<=r;k++ {
            inums[k]=tmp[k-l]
        }
}