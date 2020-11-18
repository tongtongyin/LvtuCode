package main

import "fmt"

// 给定一个数组nums，编写一个函数将所有0移动到数组的末尾，同时保持非零元素的相对顺序。
// 输入: [0,1,0,3,12] 输出: [1,3,12,0,0]

// 类似于冒泡排序的方法：将一个个的0都冒泡的最后面
func moveZeroes(nums []int) {
	n := len(nums)
	for n > 0 {
		for i,j := 0, 1; j < n; {
			if nums[i] == 0 && nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
			}
			i++
			j++
		}
		n--
	}
}

// 遍历：发现0将0删除并且追加到末尾
func moveZeroes2(nums []int) []int {
	for i, _ := range nums {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		}
	}
	return nums
}



func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)

}
