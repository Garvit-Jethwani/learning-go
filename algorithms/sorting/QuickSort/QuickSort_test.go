package QuickSort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Scenario 1: Sorting an array with positive integers",
			input:    []int{5, 2, 9, 1, 7},
			expected: []int{1, 2, 5, 7, 9},
		},
		{
			name:     "Scenario 2: Sorting an array with negative integers",
			input:    []int{-8, -3, -1, -6, -4},
			expected: []int{-8, -6, -4, -3, -1},
		},
		{
			name:     "Scenario 3: Sorting an array with mixed positive and negative integers",
			input:    []int{4, -2, 7, 0, -5, 3},
			expected: []int{-5, -2, 0, 3, 4, 7},
		},
		{
			name:     "Scenario 4: Sorting an array with duplicate elements",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
		},
		{
			name:     "Scenario 5: Sorting an empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Scenario 6: Sorting an array with a single element",
			input:    []int{42},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			quickSort(arr, 0, len(arr)-1)

			if !reflect.DeepEqual(arr, tt.expected) {
				t.Errorf("quickSort(%v) = %v, expected %v", tt.input, arr, tt.expected)
			}
		})
	}
}
