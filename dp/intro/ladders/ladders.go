// For n <= N, F(n) = F(n-1) + F(n-2) + ... + F(n-K)

package main

import "fmt"

var dp []int

// Top-Down
func ladders(n, k int) int {
	if n == 0 || n == 1 {
		return dp[0]
	}
	if dp[n] == 0 {
		for i := n - 1; i >= n-k && i >= 0; i-- {
			if dp[i] == 0 {
				dp[i] = ladders(i, k)
			}
			dp[n] += dp[i]
		}
	}
	return dp[n]
}

// Bottom-Up
func ladders2(n, k int) int {
	for i := 2; i <= n; i++ {
		for j := i - 1; j >= i-k && j >= 0; j-- {
			dp[i] += dp[j]
		}
	}
	return dp[n]
}

// Bottom-Up with Window
func ladders3(n, k int) int {
	w, cnt := 1, 2
	for i := 2; i <= n; i++ {
		if cnt <= k {
			dp[i] = w + dp[i-1]
			w = dp[i]
			cnt++
			continue
		}
		w = 2*dp[i-1] - dp[i-k-1]
		dp[i] = w
	}
	return dp[n]
}

func main() {
	n, k := 5, 3
	dp = make([]int, n+1)
	dp[0], dp[1] = 1, 1
	fmt.Println(ladders(n, k))
	dp = make([]int, n+1)
	dp[0], dp[1] = 1, 1
	fmt.Println(ladders2(n, k))
	dp = make([]int, n+1)
	dp[0], dp[1] = 1, 1
	fmt.Println(ladders3(n, k))

}
