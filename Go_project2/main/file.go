package main

import (
	"sort"
)

func findKthLargest(nums []int, k int) int {
	var uniqueSlice []int
	if k > len(nums) {
		return -1000000
	}
	for _, v := range nums {
		found := false
		for _, u := range uniqueSlice {
			if u == v {
				found = true
				break
			}
		}
		if !found {
			uniqueSlice = append(uniqueSlice, v)
		}
	}
	sort.Ints(uniqueSlice)
	for i, j := 0, len(uniqueSlice)-1; i < j; i, j = i+1, j-1 {
		uniqueSlice[i], uniqueSlice[j] = uniqueSlice[j], uniqueSlice[i]
	}

	return uniqueSlice[k-1]
	//add
}
