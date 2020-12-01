package main

// 给定一个haystack字符串和一个needle字符串，在haystack字符串中找出needle字符串出现的第一个位置 (从0开始)。如果不存在，则返回-1。
// 输入: haystack = "hleello", needle = "ll"  ->输出: 2
func strStr(haystack, needle string) int {
	i, j := 0, 0
	index := -1

	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			if index == -1 {
				index = i
			}
			if j == len(needle) -1 {
				return index
			}
			j++
		} else {

		}
		i++
	}
	return -1
}

func main() {

}
