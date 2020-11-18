package main

import (
	"fmt"
)

// 给定一个整数数组nums和一个目标值target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j := i+1; j < len(nums); j++ {
			if v + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// hash法:遍历数组
func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, val := range nums {
		if p, ok := m[target-val]; ok {
			return []int{p, i}
		}
		m[val] = i
	}
	return nil
}

func main() {
	nums := []int{9,1,2,11,7,15}
	target := 9
	fmt.Println(twoSum2(nums, target))
}
