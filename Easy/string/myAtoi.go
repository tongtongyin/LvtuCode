package main

import (
	"fmt"
	"math"
	"strings"
)

//请你来实现一个atoi函数，使其能将字符串转换成整数。
//首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。接下来的转化规则如下：
//如果第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字字符组合起来，形成一个有符号整数。
//假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成一个整数。
//该字符串在有效的整数部分之后也可能会存在多余的字符，那么这些字符可以被忽略，它们对函数不应该造成影响。
// 在任何情况下，若函数不能进行有效的转换时，请返回 0 。
// 输入: "4193 with words" ->输出: 4193

func myAtoi(str string) int {
	// 首先去掉字符串首尾的空格
	str = strings.TrimSpace(str)
	if str == "" {
		return 0
	}
	// 判断第一个字符
	var sign int
	switch str[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign = 1
	case '+':
		sign = 1
		str = str[1:]
	case '-':
		sign = -1
		str = str[1:]
	default:
		return 0
	}
	for i, v := range str {
		if v < '0' || v > '9' {
			str = str[:i]
			break
		}
	}

	sum := 0
	for _, v := range str {
		sum = sum*10 + int(v-'0')
		if sign == 1 && sum > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign == -1 && -sum < math.MinInt32 {
			return math.MinInt32
		}
	}


	return sign*sum
}


func main() {
	s := "words and 987"
	s = "4193 with words"
	s = "21474836460"
	s = " "
	s = "9223372036854775808"
	fmt.Println(myAtoi(s))

}
