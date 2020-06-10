/*
输入: head = [4,5,1,9], node = 5
输出: [4,1,9]
解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
*/

package main


//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

//最需要主要的是，删除这个节点，可以转换为将这个节点复制成后一个节点，剔除后一个节点
//                                      >--->
//                                     /     |
//1->2->4->5->6,删除节点4，可以变成 1->2->5->5->6
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main(){

}
