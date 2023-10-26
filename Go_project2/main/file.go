package main

func findKthLargest(nums []int, k int) int {
	if k < 1 || k > len(nums) {
		return -1
	}

	for i := 0; i < k; i++ {
		maxIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[maxIdx] {
				maxIdx = j
			}
		}
		nums[i], nums[maxIdx] = nums[maxIdx], nums[i]
	}

	return nums[k-1]
}
