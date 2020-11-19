package main

import "fmt"

// 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
// s = "leetcode" 返回 0;     s = "loveleetcode" 返回 2

// 哈希法：第一遍遍历字符串，将每个字符放到字典中，对应值为该字符出现的次数
// 第二遍遍历字符串，如果字符对应的值为一，则该字符为第一个不重复的字符。
func firstUniqChar(s string) int {
	m := map[int32]int{}
	for _, val := range s {
		m[val]++
	}
	for i, val := range s {
		if m[val] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	// s := "leetcode"
	// s2 := "loveleetcode"
	s3 := ""
	num := firstUniqChar(s3)
	fmt.Println(num)
}
