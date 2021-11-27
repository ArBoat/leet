func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 || k == 0 { return nil }
    stack :=[]int{}
    push :=func(i int) {
        for len(stack)>0 && nums[stack[len(stack)-1]] < nums[i] {
            stack = stack[:len(stack)-1]
        } 
        stack = append(stack, i)
    }
    for i:=0; i<k; i++{
        push(i)
    }
    res:=[]int{nums[stack[0]]}
    for i:=k; i<len(nums);i++{
        push(i)
        for stack[0]<=i-k {
            stack = stack[1:]
        }
        res = append(res, nums[stack[0]])
    }
    return res
}