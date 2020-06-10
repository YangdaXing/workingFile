/*
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
输入：["h","e","l","l","o"]
输出：["o","l","l","e","h"]
*/

package main

func reverseString(s []byte)  {
	for index := 0;index < len(s)/2;index++{
		v := s[index]
		s[index] = s[len(s)-index-1]
		s[len(s)-1-index]=v
	}
}

func main() {
	s1 := "hello"
	b2 := []byte(s1)
	reverseString(b2)
}
