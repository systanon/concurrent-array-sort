package utils

import "sort"

func Sort(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}
