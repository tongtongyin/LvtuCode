package main

import "fmt"

// 股票的最大利润
// 动态规划
func maxProfit1(prices []int) int {
	n := len(prices)

	// dp[i][0] 表示第i天手里没有股票的最大利润
	// dp[i][1] 表示第i天手里有股票的最大利润
	dp := make([][2]int, n)
	// dp[0][0] = 0  第0天手里没股票
	dp[0][1] = -prices[0]	// 第0天手里有股票
	for i := 1; i < n; i++ {
		// 第i天手里没股票的最大利润=第i-1天手里没股票 或者 第i-1天手里有股票但在第i天卖出了
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		// 第i天手里有股票的最大利润=第i-1天手里有股票 或者 第i-1天手里没股票但在第i天买入了
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
	}
	// 最后一天手里没有股票肯定是利润最大的，因为不管多少钱卖的，总比不卖强
	return dp[n-1][0]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//dp := make([][2]int, 3)
	//fmt.Println(dp)
	Profit := []int{7,1,5,3,6,4}
	res := maxProfit1(Profit)
	fmt.Println(res)

}










