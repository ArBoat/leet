func findMinArrowShots(points [][]int) int {
    if len(points) == 0 { return 0}
    sort.Slice(points, func(a, b int) bool {
        return points[a][1]<points[b][1]
    })
    count:=1
    end := points[0][1]
    for _, v:= range points{
        if v[0]> end {
            count++
            end = v[1]
        }
    }
    return count
}
/*
找到无重叠区间数量
*/