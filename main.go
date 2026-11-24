package main

import "fmt"

// MinMultiplications returns the fewest scalar multiplications needed to
// compute the product of a chain of matrices, where matrix i has dimensions
// dims[i] × dims[i+1]. Interval DP over chain segments, shortest first.
func MinMultiplications(dims []int) int {
	n := len(dims) - 1 // number of matrices
	if n < 2 {
		return 0
	}
	// dp[i][j]: cost of multiplying matrices i..j (inclusive, 0-based).
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			dp[i][j] = int(^uint(0) >> 1)
			for k := i; k < j; k++ { // last multiplication splits at k
				cost := dp[i][k] + dp[k+1][j] + dims[i]*dims[k+1]*dims[j+1]
				if cost < dp[i][j] {
					dp[i][j] = cost
				}
			}
		}
	}
	return dp[0][n-1]
}

func main() {
	// 10x20 · 20x30 · 30x40: ((AB)C) = 6000+12000 = 18000 vs (A(BC)) = 24000+8000.
	fmt.Println(MinMultiplications([]int{10, 20, 30, 40}))
}
