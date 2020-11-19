package main

import (
	"fmt"
)

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

// 输入: s = "anagram", t = "nagaram"  -> 输出: true
// 输入: s = "rat", t = "car"  -> 输出: false

// 哈希法：将字符串s做成哈希表m，key为字符，value为字符出现次数
// 遍历字符串t，将t中对应字符的出现次数在哈希表m中减去；
// 最后，遍历哈希表，如果每个value都为0则是字母异位次，否则不是
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	} else {
		m := map[int32]int{}
		for _, v := range s {
			m[v]++
		}
		for _, v := range t {
			m[v]--
		}
		for _, value := range m {
			if value != 0 {
				return false
			}
		}
	}
	return true
}

// 还可以对两个字符串进行排序，看看排序后两个字符串是否相等


func main() {
	s, t := "anagram", "nagaram"
	s1, t1 := "rat", "car"

	fmt.Println(isAnagram(s, t))
	fmt.Println(isAnagram(s1, t1))
}
