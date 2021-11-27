// 1排序

// 2投票法
func majorityElement(nums []int) int {
    count, target:=0,0
    for _, v:= range nums{
        if count == 0 {
            target = v
        }
        if target == v {
            count ++
        } else {
            count --
        }
    }
    return target
}




// 3分治

func majorityElement(nums []int) int {
    return divide(nums, 0, len(nums)-1)
}

func divide(nums []int, l ,r int) int {
    if l == r {
        return nums[l]
    }
    mid := l + (r-l)/2
    left:=divide(nums, l, mid)
    right:=divide(nums, mid+1, r)
    if left == right {
        return left
    }
    leftCount:= count(nums, l, r, left)
    rightCount:= count(nums, l, r, right)
    if leftCount > rightCount {
        return left
    } else {
        return right
    }

}

func count(nums []int, l, r, target int ) int {
    res:=0
    for i:=l;i<=r;i++ {
        if nums[i]==target{
            res++
        }
    }
    return res
}
