package main

import (
	"fmt"
	"go-sort/handlers"
	"go-sort/utils"
	"os"
	"sort"
	"sync"
)

const PARTS = 4

func main() {
	arr, err := handlers.InputHandler(PARTS, os.Stdin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	chunks := utils.Split(arr, PARTS)
	sortedParts := concurrency(chunks, sortPartition)
	var result []int
	for _, part := range sortedParts {
		result = append(result, part...)
	}
	sort.Ints(result)
	fmt.Println("Final sorted array: ", result)
}

func sortPartition(arr []int, count int) []int {
	fmt.Printf("Go Routines %d, array %v\n", count, arr)
	return utils.Sort(arr)
}

func concurrency(chunks [][]int, cb func(arr []int, count int) []int) [][]int {
	var wg sync.WaitGroup
	sortedParts := make([][]int, PARTS)

	for i := 0; i < PARTS; i++ {
		wg.Add(1)
		go func(index int, count int) {
			defer wg.Done()
			sortedParts[index] = cb(chunks[index], count)
		}(i, i+1)

	}
	wg.Wait()

	return sortedParts
}
