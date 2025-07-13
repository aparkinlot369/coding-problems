package main

import (
	"fmt"
)

var dp []int

func jumps(arr []int) int {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > 0 {
			for j := 1; j <= arr[i]; j++ {
				//fmt.Println(arr[i], i+j)
				if j+i >= len(arr) {
					break
				}
				dp[j+i] = min(dp[i]+1, dp[j+i])
				//fmt.Println(dp[j+i])
			}
		}
	}
	return dp[len(arr)-1]
}

func jumpsGreedy(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	jumps, end, lim := 0, 0, 0
	for i := 0; i < n-1; i++ {
		lim = max(lim, i+nums[i])
		if i == end {
			jumps++
			end = lim
		}
	}
	return jumps
}

func main() {
	arr := []int{3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 2, 5}
	dp = make([]int, len(arr))
	dp[0] = 0
	for i := 1; i < len(arr); i++ {
		dp[i] = 9999
	}
	fmt.Println(jumps(arr))
	fmt.Println(jumpsGreedy(arr))
}
