/*
根据每日 气温 列表，请重新生成一个列表，对应位置的输出是需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。

例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数

*/

package main

import "fmt"

/****************************方法一 暴力解法***********************************/
func dailyTemperatures(T []int) []int {
	length := len(T)
	var res []int = make([]int, length)
	for index, v := range T {
		flag := false
		for j := index + 1; j < length; j++ {
			if T[j] > v {
				res[index] = j - index
				flag = true
				break
			}
		}
		if flag != true {
			res[index] = 0
		}
	}
	return res
}

/******************************方法二*****************************************/
/*利用一个单调减栈,
1 当前温度比栈内温度低,入栈.
2 当前温度比栈内温度高,将栈内温度弹出,并记录距离差,存入res数组
3 当遍历完数组,栈内仍有,则对应res数组记录为0*/
func dailyTemperatures1(T []int) []int {
	length := len(T)
	res := make([]int, length)
	stack := make([]int, 1)
	for index, val := range T {
		for len(stack) > 0 && val > T[stack[len(stack)-1]] {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[tmp] = index - tmp
		}
		stack = append(stack, index)
	}
	return res
}

func main() {
	temperature := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperature))
	fmt.Println(dailyTemperatures1(temperature))
}
