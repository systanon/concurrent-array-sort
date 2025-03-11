package utils

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		parts    int
		expected [][]int
	}{
		{
			name:  "Even division",
			arr:   []int{1, 2, 3, 4, 5, 6, 7, 8},
			parts: 4,
			expected: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
				{7, 8},
			},
		},
		{
			name:  "Odd division",
			arr:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			parts: 4,
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11},
			},
		},
		{
			name:  "5/3",
			arr:   []int{1, 2, 3, 4, 5},
			parts: 3,
			expected: [][]int{
				{1, 2},
				{3, 4},
				{5},
			},
		},
		{
			name:  "The number of parts is greater than the length of the array",
			arr:   []int{1, 2, 3},
			parts: 5,
			expected: [][]int{
				{1},
				{2},
				{3},
				nil,
				nil,
			},
		},
		{
			name:  "One part (unchanged)",
			arr:   []int{1, 2, 3, 4, 5},
			parts: 1,
			expected: [][]int{
				{1, 2, 3, 4, 5},
			},
		},
		{
			name:  "One part (unchanged)",
			arr:   []int{1, 2, 3, 4, 5},
			parts: 2,
			expected: [][]int{
				{1, 2, 3},
				{4, 5},
			},
		},
		{
			name:  "Empty array",
			arr:   []int{},
			parts: 3,
			expected: [][]int{
				nil,
				nil,
				nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Split(tt.arr, tt.parts)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected: %v, received: %v", tt.expected, result)
			}
		})
	}
}
