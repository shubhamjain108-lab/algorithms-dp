package main

import (
	"fmt"
  "math"
)

func main() {
	var matrix = []int{1,2,3,4}
	fmt.Println("Minimum number of multiplication is: ", MatrixChainMultiplication(matrix, 4))
}

func MatrixChainMultiplication(x []int, n int) int {
	
	var i int
	// allocate composed 2d array
    a := make([][]int, n)
    for i = range a {
		a[i] = make([]int, n)
	}
	
	for i=1; i<n; i++ {
		a[i][i] = 0 
	}
	
	for l:=2; l<n; l++ {
		for i=1; i<n-l+1; i++ {
			j := i+l-1
			if (j == n){
				continue
			}
			a[i][j] = math.MaxInt64
			for k:=i; k<j; k++ {
				ans := a[i][k] + a[k+1][j] + (x[i-1] * x[k] * x[j])
				if (ans < a[i][j]) {
					a[i][j] = ans
				}
			}
		}
	}
	return a[1][n-1]
}
