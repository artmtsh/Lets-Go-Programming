package main

func findKthLargest(nums []int, k int) int {
	if k < 1 || k > len(nums) {
		return -1
	}
	maxVals := make([]int, k)
	for _, v := range nums {
		for i := 0; i < k; i++ {
			if v > maxVals[i] {
				for j := k - 1; j > i; j-- {
					maxVals[j] = maxVals[j-1]
				}
				maxVals[i] = v
				break
			}
		}
	}

	return maxVals[k-1]
}
