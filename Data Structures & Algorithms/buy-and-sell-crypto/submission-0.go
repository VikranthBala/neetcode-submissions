func maxProfit(prices []int) int {
	profit := 0
	b := 101
	for i := 0; i < len(prices); i++ {
		if prices[i] < b {
			b = prices[i]
		}
		profit = max(profit, prices[i]-b)
	}
	return profit
}