package array

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
**/

func MaxProfit(prices []int) int {
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
