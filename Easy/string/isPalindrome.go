package main

import "fmt"

// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。空字符串定义为回文串。
// 输入: "4A man, a plan, a canal: Panama" ->输出: true

func isPalindrome(s string) bool {
	sArr := []int32{}
	for _, v := range s {
		if v >= 65 && v <=90 {
			v += 32
		}
		if (v >= 48 && v <=57) || (v >= 97 && v <= 122) {
			sArr = append(sArr, v)
		}
	}
	for i, _ := range sArr {
		if sArr[i] != sArr[len(sArr)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama3"
	fmt.Println(isPalindrome(s))
}
