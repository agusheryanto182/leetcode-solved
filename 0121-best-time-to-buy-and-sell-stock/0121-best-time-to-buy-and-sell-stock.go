// func maxProfit(prices []int) int {
// 	if len(prices) == 0 {
// 		return 0
// 	}

// 	minPrice := prices[0]
// 	maxProfit := 0

// 	for i := 1; i < len(prices); i++ {
// 		if prices[i] < minPrice {
// 			minPrice = prices[i]
// 		} else {
// 			profit := prices[i] - minPrice
// 			if profit > maxProfit {
// 				maxProfit = profit
// 			}
// 		}
// 	}

// 	return maxProfit
// }

func maxProfit(prices []int) int {
    left, right, profit, result := 0, 1, 0, 0

    for right < len(prices) {
        if prices[left] < prices[right] {
            profit = prices[right] - prices[left]
            result = max(result, profit)
        } else {
            left = right
        }
        right++
    }

    return result
}