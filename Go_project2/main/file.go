package main

func findKthLargest(nums []int, k int) int {
	if k < 1 || k > len(nums) {
		return -1
	}
	for i := 1; i < k; i++ {
		maxVal := nums[i]
		maxIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > maxVal {
				maxVal = nums[j]
				maxIdx = j
			}
		}
		nums[i], nums[maxIdx] = nums[maxIdx], nums[i]
	}
	return nums[k-1]
}
