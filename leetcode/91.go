package leetcode

func Sol_medium_91(s string) int {
	n := len(s)

	if n == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && s[i-2:i] >= "10" && s[i-2:i] <= "26" {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

// 91. Decode Ways
// https://leetcode.com/problems/decode-ways/
////
// 1. state
//  - take 1 digit: '1' to '9'
//  - take 2 digits: '10' to '26'
//  => choose 1 or 2 digits => dynamic programming
// 2. define DP array
//  - dp[i] = number of ways to decode s[0:i] (the first i characters)
//  - base case: dp[0] = 1 (1 way to decode, which is do nothing)
// 3. transition
//  - if s[i-1] is valid (1-9), then dp[i] += dp[i-1]
//  - if s[i-2:i] is valid (10-26), then dp[i] += dp[i-2]
// 4. return dp[n], where n = len(s)
