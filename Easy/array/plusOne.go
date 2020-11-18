package main

import (
	"fmt"
)

// 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//你可以假设除了整数 0 之外，这个整数不会以零开头。
// 输入：digits = [1,2,3] 输出：[1,2,4] 解释：输入数组表示数字 123。
// 输入：digits = [0,0] 输出：[0,1]
// 输入：digits = [1,2,9] 输出：[1,3,0]

func plusOne(digits []int) []int {
	l := len(digits)
	// 从后往前遍历，如果数组最后一位不等于9就将数字+1,然后返回数组即可
	// 如果数字等于9，就应该将该位置变为0，继续循环将前一位+1
	for i := l-1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}else {
			digits[i] = 0
		}
	}
	// 如果数组的第0位是9，那么处理之后，总的位数会增加，在数组头多加一位1
	return append([]int{1}, digits...)
}



func main() {
	nums := []int{9,9}
	res := plusOne(nums)
	fmt.Println(res)

}
