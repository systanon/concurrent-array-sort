package utils

func SplitArray(arr []int, parts int) [][]int {
	partitions := make([][]int, parts)
	n := len(arr)
	if n == 0 {
		return partitions
	}
	if n < parts {
		parts = n
	}

	chunkSize := n / parts
	extra := n % parts
	index := 0
	for i := 0; i < parts; i++ {
		currentSize := chunkSize
		if i < extra {
			currentSize++
		}

		partitions[i] = arr[index : index+currentSize]
		index += currentSize
	}

	return partitions
}
