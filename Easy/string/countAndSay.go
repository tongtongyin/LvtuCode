package main

import (
	"strconv"
	"strings"
)

// 要描述一个数字字符串，首先要将字符串分割为最小数量的组，每个组都由连续的最多相同字符组成。
//然后对于每个组，先描述字符的数量，然后描述字符，形成一个描述组。
//要将描述转换为数字字符串，先将每组中的字符数量用数字替换，再将所有描述组连接起来。
// 11",
//"21",
//"1211",
//"111221",
//"312211",
//"13112221",
//"1113213211",
//"31131211131221",
//"13211311123113112211",
//"11131221133112132113212221",
//"3113112221232112111312211312113211",


func countAndSay(n int) string {
	str :="1"
	for i:=2;i<=n;i++{
		var tmp strings.Builder
		preByte := str[0]
		count := 1
		for j:=1;j<len(str);j++{
			if str[j]==preByte{
				count++
			}else {
				tmp.WriteString(strconv.Itoa(count))
				tmp.WriteByte(preByte)
				preByte =str[j]
				count=1
			}
		}
		tmp.WriteString(strconv.Itoa(count))
		tmp.WriteByte(preByte)
		str = tmp.String()
	}
	return str
}

func main() {
	
}
