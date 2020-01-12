package main

import (
	"fmt"
)

func main() {
	var bitonicarr = []int{1,11,2,10,4,5,2,1}
	fmt.Println("Longest bitonic subsequence :", LongestBitonic(bitonicarr, 8))
}

func LongestBitonic(arr []int, n int) int {

    lis := make([]int, n)

    lds := make([]int, n)

    for i:=0; i<n; i++ {
		lis[i] = 1
		lds[i] = 1
	}
	
	for i:=1; i<n; i++ {
		for j:=0; j<i; j++ {
			if (arr[i] > arr[j] && lis[i] < lis[j] + 1) {
				lis[i] = lis[j] + 1
			}
		}
	}

	for i:=n-2; i>=0; i-- {
		for j:=n-1; j>i; j-- {
			if (arr[i] > arr[j] && lds[i] < lds[j] + 1) {
				lds[i] = lds[j] + 1
			}
		}
	}

	max := lis[0] + lds[0] - 1

	for i:=1; i<n; i++ {
		if (lis[i] + lds[i] - 1 > max) {
			max = lis[i] + lds[i] - 1
		}
	}
    return max
}
