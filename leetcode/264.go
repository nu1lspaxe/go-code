package leetcode

func Sol_medium_264(n int) int {
	uglyNums := make([]int, n)
	uglyNums[0] = 1

	i2, i3, i5 := 0, 0, 0

	min := func(a, b, c int) int {
		if a <= b && a <= c {
			return a
		}
		if b <= a && b <= c {
			return b
		}
		return c
	}

	for i := 1; i < n; i++ {
		next2 := uglyNums[i2] * 2
		next3 := uglyNums[i3] * 3
		next5 := uglyNums[i5] * 5

		nextUgly := min(next2, next3, next5)
		uglyNums[i] = nextUgly

		if nextUgly == next2 {
			i2++
		}
		if nextUgly == next3 {
			i3++
		}
		if nextUgly == next5 {
			i5++
		}
	}

	return uglyNums[n-1]
}

// 264. Ugly Number II
// https://leetcode.com/problems/ugly-number-ii
////
// 1. An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
// 2. Use three pointers to generate ugly number list
