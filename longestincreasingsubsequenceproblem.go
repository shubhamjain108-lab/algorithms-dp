package main

import (
	"fmt"
  "math"
)

func main() {
	var arr = []int{3,4,-1,0,6,2,3}
	fmt.Println("length of longest increasing subsequnece is: ", LongestIncreasingSubsequence(arr, 7))
}
func LongestIncreasingSubsequence(x []int, n int) int {
	len := make([]int, n)
	var max int

	for i:=0; i<n; i++ {
		len[i] = 1
	}

	for i:=1; i<n; i++ {
		for j:=0; j<i; j++ {
			if(x[i] > x[j] && len[i] < len[j] + 1) {
				len[i] = len[j] + 1
			}
		}
	}
	
	for i:=0; i<n; i++ {
		if (max < len[i]) {
			max = len[i]
		}		
	}
	return max
}
