package leetcode

// 96. Unique Binary Search Trees
// https://leetcode.com/problems/unique-binary-search-trees/
////
// 1. Let dp[n] be the number of unique BSTs that can be made with n nodes
// 2. To build a BST with n nodes, consider each number i from 1 to n as the root
//  - The number less then i forms the left subtree (i.e., i-1 nodes)
//  - The number greater than i forms the right subtree (i.e., n-i nodes)
// 3. the number of BSTs with i as the root is dp[i-1] * dp[n-i]

func Sol_medium_96(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}

// Time complexity is O(n^2)

// Optimize with Catalan Formula
// https://en.wikipedia.org/wiki/Catalan_number
// dp[n] = (2n)! / ((n+1)! * n!)
// the next one is (n+1)
// (2(n+1))! / ((n+2)! * (n+1)!) = dp[n+1]
// dp[n+1] / dp[n] = (2n+2) * (2n+1) / ((n+2) * (n+1)) = 2(2n+1) / (n+2)

func Sol_medium_96_optimized(n int) int {
	res := 1

	for i := 0; i < n; i++ {
		res = res * 2 * (2*i + 1) / (i + 2)
	}

	return res
}
