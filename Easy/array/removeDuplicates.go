package main

import "fmt"

// 删除排序数组中的重复项

// 双指针法，[0,0,1,1,1,2,2,3,4]->[0,1,2,3,4,2,2,3,4]
func removeDuplicates1(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

//用切片删除多余元素法， [0,0,1,1,1,2,2,3,4]->[0,1,2,3,4]
func removeDuplicates2(nums []int) int {
	// 从后往前遍历
	for i := len(nums)-1; i > 0; i-- {
		if nums[i] != nums[i-1] {
			continue
		}else {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}

func main() {
	var nums = []int{0,0,1,1,1,2,2,3,4}
	l := removeDuplicates2(nums)
	fmt.Println(l)
}
