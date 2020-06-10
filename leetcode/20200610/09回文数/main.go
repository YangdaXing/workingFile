/*判断一个整数是否是回文数。回文数是指
正序（从左向右）和倒序（从右向左）读都是一样的整数。
输入：1212
返回false

进阶：不使用整数转字符串
*/

package main

import (
	"fmt"
	"strconv"
)

//将整数转换成字符串进行比较
func isPalindrome(x int) bool {
	tmp := strconv.Itoa(x)
	lenth := len(tmp)
	var(
		i int =0
		j int = lenth-1
	)
	for ;i<=j; {
		if tmp[i]!=tmp[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindrome2(x int) bool{
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x<0 || (x%10 == 0 && x !=0){
		return false
	}
	tmp := 0
	for x > tmp {
		tmp = tmp*10 + (x % 10)
		x /= 10
	}
	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。

	return x == tmp || x == (tmp/10)
}

func main(){
	a := 12321
	fmt.Println(isPalindrome2(a))
}
