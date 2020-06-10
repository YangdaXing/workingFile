package main

/*
给定一棵二叉搜索树，请找出其中第k大的节点。
示例 1:

输入: root = [3,1,4,null,2], k = 1
      3
     / \
    1   4
         \
          2
输出: 4
*/

var tmp []int
func kthLargest(root *TreeNode, k int) int {
	InOrderTravel(root)
	return tmp[k-1]
}

func InOrderTravel(root *TreeNode){
	if root == nil {
		return
	}
	InOrderTravel(root.Left)
	tmp = append(tmp,root.Val)
	InOrderTravel(root.Right)
}


func main() {
	
}
