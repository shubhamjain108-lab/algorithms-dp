package main

import (
	"fmt"
	"math"
)

func main() {
	str1 := "AGGTAB"
	str2 := "GXTXAYB"
	len1 := len(str1)
	len2 := len(str2)
	fmt.Println("length of shortest common supersequnece is: ", ShortestCommonSupersequence([]byte(str1), []byte(str2), len1, len2))
}

func LongestCommonSubsequence(x []byte, y []byte, n int, m int) float64 {

	// allocate composed 2d array
    a := make([][]float64, n+1)
    for i := range a {
		a[i] = make([]float64, m+1)
    }

	for i := 0; i<=n; i++ {
		for j := 0; j<=m; j++ {
			if (i == 0 || j == 0){
				a[i][j] = 0
			} else if (x[i - 1] == y[j - 1]){
				a[i][j] = a[i-1][j-1] + 1
			} else {
				a[i][j] = math.Max(a[i-1][j], a[i][j-1])
			}
		}
	}
	return a[n][m]
}

func ShortestCommonSupersequence(x []byte, y []byte, n int, m int) float64 {
	len := LongestCommonSubsequence(x, y, n, m)
	return float64(n + m) - len
}
