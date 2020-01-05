package main

import (
	"fmt"
)

func main() {
	var profit = []int{2,3,1,4}
	var weight = []int{3,4,6,5}
	var W int = 8
	fmt.Println(KnapSack(profit, weight, W,  4))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func KnapSack(profit []int, weight []int, W int, n int) int {
	
	// allocate composed 2d array
    sol := make([][]int, n+1)
    for i := range sol {
		sol[i] = make([]int, W+1)
	}
	
	for i:=0; i<=n; i++ {
		for j:=0; j<=W; j++ {
			if (i == 0 || j == 0) {
				sol[i][j] =  0
			} else if (weight[i-1] <= j) { 
				sol[i][j] = max(profit[i-1] + sol[i-1][j-weight[i-1]], sol[i-1][j])
			} else {
				sol[i][j] = sol[i-1][j]
			}
		}
	}
	return sol[n][W]
}
