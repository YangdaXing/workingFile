/*请完成一个函数，输入一个二叉树，该函数输出它的镜像。
例如输入：
     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：
     4
   /   \
  7     2
 / \   / \
9   6 3   1

示例 1：

输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]
*/

package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	//首先判断根节点是否为nil
	if root == nil {
		return nil
    }
    tmp := root.Left
    root.Left = mirrorTree(root.Right)
    root.Right = mirrorTree(tmp)
    return  root
}



func main() {
	
}
