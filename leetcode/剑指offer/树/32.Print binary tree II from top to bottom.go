package main

/*
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
例如:
给定二叉树: [3,9,20,null,null,15,7],
    3
   / \
  9  20
 /     \
15      7
返回其层次遍历结果：
[
  [3],
  [9,20],
  [15,7]
]
*/

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	//利用队列
	queue := []*TreeNode{root}
	var level int = 0
	for len(queue) != 0 {
		//利用临时队列，暂存每个节点的左右子树
		temp := []*TreeNode{}
		//只需考虑在同一层上追加元素
		res = append(res,[]int{})
		//遍历队列，追加队列元素到切片同一层，追加队列元素左右子树到临时队列
		for _,v := range queue {
			res[level] = append(res[level],v.Val)
			if v.Left != nil {
				temp = append(temp,v.Left)
			}
			if v.Right != nil{
				temp = append(temp,v.Right)
			}
		}
		//层级加1，队列重新复制为队列的左右子树集
		level++
		//队列赋值
		queue = temp
	}
	return res
}

func main() {
	
}
