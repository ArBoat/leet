func findDuplicate(nums []int) int {
    if len(nums) == 0 { return 0 }
    slow, fast:= 0, 0
    slow = nums[slow]
    fast = nums[nums[fast]]
    for slow != fast {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }
    cur:=0
    for cur != slow {
        cur = nums[cur]
        slow = nums[slow]
    }
    // IM: return cur
    return cur
}

/*
O(n)/O(1)
转化成142循环链表II,秒啊
*/



func findDuplicate(nums []int) int {
    l, r := 1, len(nums)-1
    res:=-1
    for l<=r {
        mid:=l+(r-l)/2
        count:=0
        for _,v:=range nums {
            if mid >= v{
                count++
            }
        }
        if count > mid {
            r = mid - 1
            // 重要
            res = mid
        } else {
            l = mid + 1
        } 
    }
    return res
}

/*
O(nlogn)/O(1)
抽屉原理 + 二分查找
*/