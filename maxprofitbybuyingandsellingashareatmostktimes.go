package main

import (
	"fmt"
)

func main() {
	var price = []int{10,22,5,75,65,80}
	var k int = 2
	fmt.Println("maximum profit by buying and selling a share at most kth times:", dp.MaxProfit(price, k, 6))
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MaxProfit(price []int, k int, n int) int {
	
	profit := make([][]int, k+1)
	for i := range profit {
		profit[i] = make([]int, n)
	}

	for i:=0; i<k; i++ {
		profit[i][0] = 0
	}

	for j:=0; j<n; j++ {
		profit[0][j] = 0
	}

	for i:=1; i<=k; i++ {
		lastval := math.MinInt64
		for j:=1; j<n; j++ {
			lastval = Max(lastval, profit[i - 1][j - 1] - price[j - 1])
			profit[i][j] = Max(profit[i][j - 1], price[j] + lastval)
		}
	}
    return profit[k][n - 1]
}

