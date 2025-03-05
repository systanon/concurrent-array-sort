package main

import (
	"bufio"
	"fmt"
	"go-sort/utils"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const PARTS = 4

func main() {

	fmt.Println("Enter a series of integers separated by spaces:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	strNums := strings.Fields(input)

	var arr []int
	for _, str := range strNums {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Invalid input. Please enter integers only.")
			return
		}
		arr = append(arr, num)
	}
	if len(arr) < PARTS {
		fmt.Printf("Please enter minimum %d integers\n", PARTS)
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
	sort.Ints(arr)
	sortedParts <- arr
}
