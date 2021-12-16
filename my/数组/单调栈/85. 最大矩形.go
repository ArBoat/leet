func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 { return 0 }
    heights:=make([]int, len(matrix[0]))
    res:=0
    for _, row:= range matrix {
        for j, column := range row {
            if column == '1' {
                heights[j]+=1
            } else {
                heights[j]=0
            }
        }
        res = max(res, lagest(heights))
    }
    return res
}

func lagest(heights[]int) int {
    if len(heights) == 0 { return 0 }
    if len(heights) == 1 { return heights[0]}
    res:=0
    stack, newHeights:=[]int{}, []int{0}
    newHeights = append(newHeights, heights...)
    newHeights = append(newHeights,0)
    for i, v:= range newHeights {
        for len(stack)>0 && newHeights[stack[len(stack)-1]] > v {
            left:= stack[len(stack)-2]
            h:= stack[len(stack)-1]
            res = max(res, (i-left-1)*newHeights[h])
            // *******千万别忘记弹出************
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    return res
}

func max(a, b int) int {
    if a > b { 
        return a
    }
    return b
}

/*
O(mn) /O(mn)
84题pulus
*/