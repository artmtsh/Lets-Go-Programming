package main

import "fmt"

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

func main() {
	nums := []int{4, 3, 6, 4, 1}
	k := 3
	fmt.Println(findKthLargest(nums, k))
}
