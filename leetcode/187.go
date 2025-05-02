package leetcode

func Sol_medium_187(s string) []string {
	seen := make(map[string]int)
	res := []string{}

	n := len(s)
	if n < 10 {
		return res
	}

	for i := 0; i <= n-10; i++ {
		sub := s[i : i+10]
		seen[sub]++
		if seen[sub] == 2 {
			res = append(res, sub)
		}
	}
	return res
}

// 187. Repeated DNA Sequences
// https://leetcode.com/problems/repeated-dna-sequences
////
// 1. Constraints
//  - Only substrings of exactly length 10
//  - A substring can overlap with others -> sliding window
//  - Only return the substrings that occur > 1 time
// 2. Steps
//  - Use a map to record all 10-letter substrings
//  - Move a sliding window of size 10 across the string
//  - For each 10-letter substring, record how many times it appears
// 3. Return the substrings that appear more than once
