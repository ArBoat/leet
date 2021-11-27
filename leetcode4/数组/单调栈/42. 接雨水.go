func trap(height []int) int {
    l, r:= 0, len(height)-1
    lmax, rmax, res := 0, 0, 0
    for l<r{
        if height[l] < height[r] {
            if height[l] >= lmax {
                lmax = height[l]
            } else {
                res+=lmax-height[l]
            }
            //重要
            l++
        } else {
            if height[r]>=rmax {
                rmax = height[r]
            } else {
                res+= rmax - height[r]
            }
            // 重要
            r--
        }
    }
    return res
}

/*

*/