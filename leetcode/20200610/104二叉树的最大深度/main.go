package main


//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/**************************方法一************************************/

func max(a,b int) int {
    if a > b{
        return a
    }
    return b
}

func maxDepth(root *TreeNode) int {
   if root == nil {
       return 0
   }

   left := root.Left
   right := root.Right
   return max(maxDepth(left),maxDepth(right))+1
}

/**************************方法二************************************/
var mDeep int
func maxDepth1(root *TreeNode) int {
    mDeep = 0
    if root != nil{
        recursive(root,1)
    }
    return mDeep
}

func recursive(node *TreeNode,deep int){
    if node.Left != nil{
        recursive(node.Left,deep + 1)
    }
    if node.Right != nil{
        recursive(node.Right,deep + 1)
    }
    //比较两个分支的最大长度，保存在全局变量mDeep
    if node.Left == nil && node.Right == nil{
        if deep > mDeep {
            mDeep = deep
        }
    }
}

func main() {


}
