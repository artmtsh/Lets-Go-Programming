package main

func findKthLargest(nums []int, k int) int {
	if k < 1 || k > len(nums) {
		return -1
	}
	k--
	for {
		pivotIndex := partition(nums)
		if k == pivotIndex {
			return nums[k]
		} else if k < pivotIndex {
			nums = nums[:pivotIndex]
		} else {
			nums = nums[pivotIndex+1:]
			k -= (pivotIndex + 1)
		}
	}
}

func partition(nums []int) int {
	pivotIndex := len(nums) - 1
	pivot := nums[pivotIndex]
	left := 0
	for i := 0; i < pivotIndex; i++ {
		if nums[i] > pivot {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
	nums[left], nums[pivotIndex] = nums[pivotIndex], nums[left]
	return left
}
