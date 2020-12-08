package array

import "fmt"

/** 122
* 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
* 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
* 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*
* 输入: [7,1,5,3,6,4]
* 输出: 7
*
* 分析
* 购买股票时，手中需要没有股票
* 动态规划
* 定义dp[i][j],i为天数j为持股状态
* dp[i][0]=max{dp[i−1][0],dp[i−1][1]+prices[i]}
* dp[i][1]=max{dp[i−1][0]-prices[i],dp[i−1][1]}
**/
//动态规划 时间复杂度O(N) 空间复杂度O(1)
func MaxProfit(prices []int) int {
	lenPrices := len(prices)
	if lenPrices < 2 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for tmpK := 1; tmpK < lenPrices; tmpK++ {
		//不持有股票 计算未持有股票和前一天持有股票今天卖掉那个利润大
		dp[tmpK][0] = max(dp[tmpK-1][0], dp[tmpK-1][1]+prices[tmpK])
		//持有股票 计算前一天未持有股票今天购买和继续持有哪个利润大
		dp[tmpK][1] = max(dp[tmpK-1][0]-prices[tmpK], dp[tmpK-1][1])
	}
	fmt.Println(dp)
	return dp[lenPrices-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//贪心 时间复杂度O(N) 空间复杂度O(1)
func maxProfit(prices []int) int {
	lenPrices := len(prices)
	if lenPrices < 2 {
		return 0
	}
	tmpPrice := 0
	for tmpK := 1; tmpK < lenPrices; tmpK++ {
		if diff := prices[tmpK] - prices[tmpK-1]; diff > 0 {
			tmpPrice += diff
		}
	}
	return tmpPrice
}