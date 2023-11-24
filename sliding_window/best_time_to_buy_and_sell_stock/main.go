package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maximumProfit := 0
	l := 0

	for i, price := range prices {
		if price < prices[l] {
			l = i
		}
		currentProfit := price - prices[l]
		maximumProfit = max(maximumProfit, currentProfit)
	}

	return maximumProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
