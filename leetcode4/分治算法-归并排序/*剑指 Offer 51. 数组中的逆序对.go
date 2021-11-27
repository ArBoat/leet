func reversePairs(nums []int) int {
    return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) int {
   if l >= r {
       return 0
   }
   mid:=l + (r-l)/2
   count:= mergeSort(nums,l,mid) + mergeSort(nums, mid+1, r)
   j:=mid+1
   for i:=l; i<= mid; i++ {
       for j<= r && nums[i] > nums[j]{
           j++
       }
       count += j - (mid+1)
   }
   tmp :=[]int{}
   p1, p2:= l, mid+1
   for p1 <= mid && p2 <= r {
       if nums[p1] <= nums[p2] {
           tmp = append(tmp, nums[p1])
           p1++
       } else {
           tmp = append(tmp, nums[p2])
           p2++
       }
   } 
   for ;p1<=mid;p1++ {
       tmp = append(tmp, nums[p1])
   }
   for ;p2<=r;p2++ {
       tmp = append(tmp, nums[p2])
   }
   for k:=l;k<=r;k++{
       nums[k] = tmp[k-l]
   }
   return count
}