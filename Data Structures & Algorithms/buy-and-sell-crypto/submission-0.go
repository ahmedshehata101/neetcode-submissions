func maxProfit(prices []int) int {
if len(prices) == 0 {
		return 0
	}

	left := 0 // buy pointer
	maxProfit := 0

	for right := 1; right < len(prices); right++ {
		if prices[right] < prices[left] {
			// Found a better buying price
			left = right
		} else {
			// Valid selling opportunity
			profit := prices[right] - prices[left]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
