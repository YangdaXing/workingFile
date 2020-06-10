/*给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
     0
    / \
  -3   9
  /   /
-10  5
*/

package main

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return sortedBST(nums, 0, len(nums)-1)
}

func sortedBST(nums []int ,start,end int) *TreeNode{
	if start > end {
		return nil
	}
	if start == end {
		return &TreeNode{Val: nums[start]}
	}
	/*求中点不要用 int mid = (l + r)/2，有溢出风险，
	比如 l,r都是IntMax,相加会溢出
	稳妥的方法是 int mid = l + (r-l)/2;
	如果你把除2改成右移1位，会和面试官更搭哦。*/
	mid := start+((end-start)>>1)
	root := &TreeNode{Val: mid}
	root.Left = sortedBST(nums,start,mid-1)
	root.Right = sortedBST(nums,mid+1,end)
	return root

}


