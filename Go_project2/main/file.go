package main

import (
	"sort"
)

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums[k-1]
	//add
}
