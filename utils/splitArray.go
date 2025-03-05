package utils

func Split(arr []int, parts int) [][]int {
	partitions := make([][]int, parts)
	n := len(arr)
	if n == 0 {
		return partitions
	}
	if n < parts {
		parts = n
	}

	index := 0
	for i := 0; i < parts; i++ {
		remaining := n - index
		size := (remaining + (parts - i) - 1) / (parts - i)

		partitions[i] = arr[index : index+size]
		index += size
	}

	return partitions
}
