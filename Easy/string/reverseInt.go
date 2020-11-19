package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if x > 0 {
		return reverseInt(x)
	}else {
		return -reverseInt(-x)
	}
}

func reverseInt(x int) int {
	var res []int
	for x > 0 {
		num := x % 10
		x /= 10
		res = append(res, num)
	}
	var sum int
	for i := len(res)-1; i >= 0; i-- {
		sum += res[i]*int(math.Pow10(len(res)-i-1))
	}
	if sum < -int(math.Pow(2, 31)) || sum > int(math.Pow(2, 31))-1 {
		return 0
	}else {
		return sum
	}
}

func main() {
	var sum int
	sum = reverse(1534236469)
	fmt.Println(sum)
	fmt.Println(int(math.Pow(2, 31))-1)
}
