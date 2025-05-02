package leetcode

// Time Limit Exceeded
// type NumArray struct {
// 	arr []int
// }

// func Constructor(nums []int) NumArray {
// 	return NumArray{arr: nums}
// }

// func (this *NumArray) Update(index int, val int) {
// 	this.arr[index] = val
// }

// func (this *NumArray) SumRange(left int, right int) int {
// 	result := 0
// 	for i := left; i < right+1; i++ {
// 		result += this.arr[i]
// 	}
// 	return result
// }

// Optimize 1 - use a slice to store prefix sum
// For each update, the time complexity is O(n)
// type NumArray struct {
// 	arr       []int
// 	prefixSum []int
// }

// func Constructor(nums []int) NumArray {
// 	prefixSum := make([]int, len(nums)+1)
// 	for i, num := range nums {
// 		prefixSum[i+1] = prefixSum[i] + num
// 	}
// 	return NumArray{arr: nums, prefixSum: prefixSum}
// }

// func (this *NumArray) Update(index int, val int) {
// 	diff := val - this.arr[index]
// 	this.arr[index] = val

// 	for i := index + 1; i < len(this.prefixSum); i++ {
// 		this.prefixSum[i] += diff
// 	}
// }

// func (this *NumArray) SumRange(left int, right int) int {
// 	return this.prefixSum[right+1] - this.prefixSum[left]
// }

// Optimize 2 - use segment tree
type NumArray struct {
	n       int // length of array
	segTree []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	segTree := make([]int, n*2)
	// 1. Store the nodes to the leaves
	for i, num := range nums {
		segTree[n+i] = num
	}
	// 2. Sum up the nodes
	for i := n - 1; i > 0; i-- {
		segTree[i] = segTree[i*2] + segTree[i*2+1]
	}
	return NumArray{
		n:       n,
		segTree: segTree,
	}
}

func (this *NumArray) Update(index int, val int) {
	index += this.n
	this.segTree[index] = val

	for index > 1 {
		index /= 2
		this.segTree[index] = this.segTree[index*2] + this.segTree[index*2+1]
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	left += this.n
	right += this.n
	sum := 0

	for left <= right {
		if left%2 == 1 {
			sum += this.segTree[left]
			left++
		}
		if right%2 == 0 {
			sum += this.segTree[right]
			right--
		}

		left /= 2
		right /= 2
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

// 307. Range Sum Query - Mutable
// https://leetcode.com/problems/range-sum-query-mutable/
