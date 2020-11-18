package main

import "fmt"

// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
// 输入: [1,2,3,4,5,6,7] 和 k = 3
// 输出: [5,6,7,1,2,3,4]

// append拼接法:将后k个元素放到前面
func rotate1(nums []int, k int)  {
	n := len(nums)
	k = k % n
	copy(nums, append(nums[n-k:], nums[0:n-k]...))
}

// 三次翻转法：先全部旋转，然后分别旋转前k个元素，和k到最后的元素
func rotate2(nums []int, k int)  {
	n := len(nums)
	k = k % n
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
func reverse(nums []int)  {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

func main()  {
	nums := []int{1,2,3,4,5,6,7}
	k := 3
	rotate1(nums, k)
	fmt.Println(nums)
}
