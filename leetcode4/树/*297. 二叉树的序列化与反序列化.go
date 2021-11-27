/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//辅助
type Codec struct {
    res []string
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return "#"
    }
    // int 转string
    return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) +
		"," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    this.res = strings.Split(data,",")
    return this.dfs()
}

func (this *Codec) dfs() *TreeNode {
    if this.res == nil || len(this.res) == 0 {
        return nil
    }
    node := this.res[0]
    this.res = this.res[1:] // 一定在这
    if node == "#" {
        return nil
    }
    // string转int
    v, _ := strconv.Atoi(node)
    root:=&TreeNode{
        Val: v,
        Left: this.dfs(),
        Right: this.dfs(),
    }
    return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */