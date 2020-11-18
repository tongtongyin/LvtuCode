package main

import "fmt"

// 只出现一次的数字：
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 输入: [4,1,2,1,2] 输出: 4

// 字典法：每个数当做key，每个数出现的次数作为value，找出value=1的key
func singleNumber(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}
	for key, value := range m {
		if value == 1 {
			return key
		}
	}
	return 0
}

// 亦或法: 0^a=a; a^a=0; a^b^a=b^a^a=b^(a^a)=b^0=b
func singleNumber2(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

func main() {
	nums := []int{4,1,2,1,2}
	res := singleNumber2(nums)
	fmt.Println(res)

}
