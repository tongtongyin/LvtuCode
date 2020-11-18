package main

import (
	"fmt"
	"sort"
)
// 检查数组中有没有重复的数，有->true,没有->false


// 字典统计方法
func containsDuplicate(nums []int) bool {
	static := make(map[int]int)
	for _, v := range nums {
		static[v]++
		if static[v] > 1 {
			return true
		}
	}
	return false
}

// 同上一种方法，只是写法不同
func containsDuplicate2(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = 1
		}
	}
	return false
}

// 排序方法：首先对数组进行排序，然后遍历排序后的数组，如果有连着两个元素相同->true，否则->false
func containsDuplicate3(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1,2,3,4,5,6}
	nums2 := []int{1,2,3,4,5,1,6}
	a := containsDuplicate(nums)
	b := containsDuplicate(nums2)
	fmt.Println(a, b)
}
