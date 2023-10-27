package main

import "math"

func findKthLargest(nums []int, k int) int {
	if k < 1 || k > len(nums) {
		return -1
	}
	kthLargest := math.MinInt
	for k > 0 {
		currentMax := math.MinInt
		maxIndex := -1
		for i, num := range nums {
			if num > currentMax {
				currentMax = num
				maxIndex = i
			}
		}
		if maxIndex >= 0 {
			nums = append(nums[:maxIndex], nums[maxIndex+1:]...)
			kthLargest = currentMax
			k--
		} else {
			break
		}
	}
	return kthLargest
}
