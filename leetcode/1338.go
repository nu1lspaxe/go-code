package leetcode

import "sort"

func Sol_medium_1338(arr []int) int {
	n := len(arr)
	half := n / 2

	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}

	frequencies := make([]int, 0, len(freq))
	for _, f := range freq {
		frequencies = append(frequencies, f)
	}

	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i] > frequencies[j]
	})

	removed := 0
	count := 0
	for _, f := range frequencies {
		removed += f
		count++

		if removed >= half {
			break
		}
	}
	return count
}

// 1338. Reduce Array Size to The Half
// https://leetcode.com/problems/reduce-array-size-to-the-half
////
// 1. Observation
//  - Numbers that apprear more frequently are more valuable to remove
//  - Greedy : always remove the number that occurs the most first
// 2. Steps
//  - Count frequency : use a map to count how many times each number appears
//  - Sort the frequency : bigger frequency first, because biggers remove more
//  - Remove numbers :
//    - Start summing frequencies
//    - Keep a counter for how many numbers you removed
//    - Stop when the total removed elements is at least half of len(array)
// 3. Return the number of different integers picked
