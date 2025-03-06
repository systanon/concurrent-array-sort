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
	var wg sync.WaitGroup
	sortedParts := make(chan []int, PARTS)

	for i := 0; i < PARTS; i++ {
		wg.Add(1)
		go sortPartition(chunks[i], &wg, sortedParts)

	}
	wg.Wait()
	close(sortedParts)

	var result []int
	for part := range sortedParts {
		result = append(result, part...)
	}
	sort.Ints(result)
	fmt.Println("Final sorted array: ", result)
}

func sortPartition(arr []int, wg *sync.WaitGroup, sortedParts chan<- []int) {
	defer wg.Done()
	sortedParts <- arr
}
