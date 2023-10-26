package main

func findKthLargest(nums []int, k int) int {
	if k < 0 || k > len(nums) {
		return -1
	}
	maxVals := make([]int, k)
	for _, v := range nums {
		if v >= maxVals[k-1] {
			for j := 0; j < k-1; j++ {
				maxVals[j] = maxVals[j+1]
			}
			maxVals[k-1] = v
		}
	}
	return maxVals[0]
}
