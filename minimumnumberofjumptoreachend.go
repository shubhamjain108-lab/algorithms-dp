package main

import (
	"fmt"
)

func main() {
	var jumparr = []int{2,1,3,2,4,5,1,2,8}
	fmt.Println("Minimum number of jump to reach end:", MinimumJump(jumparr, 9))
}

func Min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func MinimumJump(arr []int, n int) int {
    minimum_jump_arr := make([]int, n)
	
	minimum_jump_arr[0] = 0

	for i:=1; i<n; i++ {
		minimum_jump_arr[i] = math.MaxInt64
		for j:=0; j<i; j++ {
			if (i <= j + arr[j]) {
				minimum_jump_arr[i] = Min(minimum_jump_arr[i], minimum_jump_arr[j] + 1)
			}
		}
	}
    return minimum_jump_arr[n - 1]
}
