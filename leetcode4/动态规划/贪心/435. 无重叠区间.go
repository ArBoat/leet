func eraseOverlapIntervals(intervals [][]int) int {
    if len(intervals) == 0 { return 0}
    sort.Slice(intervals, func(a, b int)bool {
        // IM:[1]
       return intervals[a][1]<intervals[b][1]
    })
    count:=1
    // IM: intervals[0][1], v[0], v[1]
    end:=intervals[0][1]
    for _, v:= range intervals {
        if v[0]>=end {
            count++
            end = v[1]
        }
    }
    return len(intervals) - count
}
/*
重叠子区间， 边界根据题意
*/