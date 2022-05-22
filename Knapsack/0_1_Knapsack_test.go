package knapsack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Problem Statement#
// Given two integer arrays to represent weights and profits of ‘N’ items, we need to find a subset of these items which will give
// us maximum profit such that their cumulative weight is not more than a given number ‘C’. Write a function that returns the maximum
// profit. Each item can only be selected once, which means either we put an item in the knapsack or skip it.

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// Time complexity O(N * C) where n is the number of items and c is the capacity of the knapsack. Space complexity O(n * c) for the 2d array
func solveKnapsack(profits, weights []int, capacity int) int {
	if capacity <= 0 || len(profits) == 0 || len(weights) != len(profits) {
		return 0
	}

	n := len(profits)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for c := 0; c <= capacity; c++ {
		if weights[0] <= c {
			dp[0][c] = profits[0]
		}
	}

	for i := 1; i < n; i++ {
		for c := 0; c <= capacity; c++ {
			if weights[i] <= c {
				dp[i][c] = max(dp[i-1][c], profits[i]+dp[i-1][c-weights[i]])
			} else {
				dp[i][c] = dp[i-1][c]
			}
		}
	}

	for i := range dp {
		fmt.Println(dp[i])
	}

	return dp[n-1][capacity]
}

func TestSolveKnapsack_Example1(t *testing.T) {
	assert.Equal(t, 22, solveKnapsack([]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 7))
}

func TestSolveKnapsack_Example2(t *testing.T) {
	assert.Equal(t, 17, solveKnapsack([]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 6))
}
