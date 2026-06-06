package main

import "fmt"

func main() {
	//n := 5
	//var d int
	//arr := make([]int, n+1)
	//for i := 0; i <= n; i++ {
	//	for j := i; j != 0; j /= 2 {
	//		d = j % 2
	//		arr[i] += d
	//	}
	//}
	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}

	miniCostPaths()
}

func miniCost() {
	n := 10
	//k := 2
	costs := []int{1, 2, 4, 6, 7, 3, 1, 2, 3, 5}
	dp := make([]int, n)
	dp[0] = 1
	mini := dp[0]
	for i := 1; i < n; i++ {
		if mini >= dp[i-1] {
			mini = dp[i-1]
		}
		dp[i] = costs[i] + mini
	}
	for i := 0; i < n; i++ {
		fmt.Println(dp[i])
	}
}

func miniCostPaths() {
	n := 10
	k := 2
	costs := []int{1, 2, 4, 6, 7, 3, 1, 2, 3, 5}
	dp := make([]int, n)
	dp1 := make([]int, n)
	dp[0] = 1
	dp1[0] = 1
	for i := 1; i < n; i++ {
		mini := dp[i-1]
		for j := 1; j <= k; j++ {
			if i-j < 0 {
				continue
			}
			if mini >= dp[i-j] {
				mini = dp[i-j]
			}
		}
		for j := 1; j <= k; j++ {
			if i-j < 0 {
				continue
			}
			if mini == dp[i-j] {
				dp1[i] += 1
			}
		}
		dp[i] = costs[i] + mini
	}
	for i := 0; i < n; i++ {
		fmt.Println(dp1[i])
	}
}
