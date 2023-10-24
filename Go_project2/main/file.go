package main

import "sort"

// Example of tested file
func findKthLargest(nums []int, k int) int {
	//your code should be there
	sort.Ints(nums)
	return nums[k]
	//add
}
