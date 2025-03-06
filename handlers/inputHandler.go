package handlers

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func InputHandler(parts int, input io.Reader) ([]int, error) {
	fmt.Println("Enter a series of integers separated by spaces:")
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	strNums := strings.Fields(scanner.Text())

	var arr []int
	for _, str := range strNums {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, errors.New("please enter integers only")
		}
		arr = append(arr, num)
	}
	if len(arr) < parts {
		fmt.Printf("Please enter minimum %d integers\n", parts)
		return nil, fmt.Errorf("please enter minimum %d integers\n", parts)
	}
	return arr, nil
}
