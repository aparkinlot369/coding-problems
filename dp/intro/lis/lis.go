package main

import (
	"fmt"
)

var dp []int

func lis(arr []int) int {
	dp = make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = max(dp[i], dp[j]+1)
				res = max(res, dp[i])
			}
		}
	}
	return res
}

func upperBound(arr []int, x int) int {
	low := 0
	high := len(arr)
	ans := len(arr)

	for low < high {
		mid := low + (high-low)/2
		if arr[mid] > x {
			ans = mid  // This could be our answer
			high = mid // Try to find an even smaller index
		} else {
			low = mid + 1 // Element is not greater, search in the right half
		}
	}
	return ans
}

func lis2(arr []int) int {
	dp = make([]int, len(arr)+1)
	inf := 1000000000
	dp[0] = -inf
	for i := 1; i < len(arr); i++ {
		dp[i] = inf
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		for j := 1; j <= len(arr); j++ {
			if dp[j-1] < arr[i] && arr[i] < dp[j] {
				dp[j] = arr[i]
				res = max(res, j)
			}
		}
	}
	return res
}

func lis3(arr []int) int {
	dp = make([]int, len(arr)+1)
	inf := 1000000000
	dp[0] = -inf
	for i := 1; i < len(arr); i++ {
		dp[i] = inf
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		l := upperBound(dp, arr[i])
		if l < len(dp) && dp[l-1] < arr[i] && arr[i] < dp[l] {
			dp[l] = arr[i]
			res = max(res, l)
		}
	}
	return res
}

func main() {
	arr := []int{50, 4, 10, 7, 9, 30, 100}
	fmt.Println(lis3(arr))
	arr = []int{3, 4, 5, 6, 20}
	fmt.Println(lis3(arr))

}
