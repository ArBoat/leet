func canJump(nums []int) bool {
    distance:=0
    for i:=0; i< len(nums)-1; i++ {
        distance = max(distance, i+nums[i])
        if distance <= i {
            return false
        }
    }
    return distance>=len(nums)-1
}

func max(a, b int) int {
    if a>b {
        return a
    }
    return b
}