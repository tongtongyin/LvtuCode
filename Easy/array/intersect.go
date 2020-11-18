package main

import (
	"fmt"
	"sort"
)

// 给定两个数组，编写一个函数来计算它们的交集。
// nums1 = [4,9,4,5], nums2 = [9,4,9,8,4] return [9,4,4]

// 哈希表法：将比较短的数组中每个数作为key，出现次数作为value
// 遍历第二个数组，将在字典中出现过的数加到新数组，并且将字典中该数对应的出现次数减一
func intersect(nums1, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}
	var res []int
	for _, p := range nums2 {
		if m[p] > 0 {
			res = append(res, p)
		}
		m[p]--
	}
	return res
}

// 排序+双指针法：排序两个数组，用两个指针指向两个数组的头
// 如果相等就将值加入结果数组，并将两指针均向后移一位；如果nums1大就将nums2的指针向后移一位，反之
func intersect2(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var res []int
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] > nums2[j] {
			j++
		}else if nums1[i] < nums2[j]{
			i++
		}else {
			res = append(res, nums1[i])
			i++
			j++
		}
	}
	return res
}


func main() {
	nums1 := []int{4,9,4,5}
	nums2 := []int{9,4,9,8,4}
	res := intersect2(nums1, nums2)
	fmt.Println(res)
}
