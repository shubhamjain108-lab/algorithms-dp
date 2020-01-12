package main

import (
	"fmt"
)

func main() {
	var substring1 = "abcdxyz"
	var substring2 = "xyzabcd"
	fmt.Println("Longest common substring :", LongestCommonSubString([]byte(substring1), []byte(substring2), 7, 7))
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func LongestCommonSubString(x, y []byte, n, m int) int {

    sol := make([][]int, n+1)
    for i:= range sol {
        sol[i] = make([]int, m+1)
    }

    max := 0

    for i:=0; i<=n; i++ {
        for j:=0; j<=m; j++ {
            if (i == 0 || j == 0){
                sol[i][j] = 0
            } else if (x[i - 1] == y[j - 1]) {
                sol[i][j] = sol[i - 1][j -1] + 1
                max = Max(max, sol[i][j])
            } else {
                sol[i][j] = 0
            }
        }
    }
    return max
}
