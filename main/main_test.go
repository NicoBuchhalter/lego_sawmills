package main

import "testing"

func TestMain_calculateIncome(t *testing.T) {
	tests := []struct {
		name     string
		trunks   []int
		expected int
	}{
		{
			name:     "Just a few small trunks",
			trunks:   []int{1, 2, 1},
			expected: 3,
		},
		{
			name:     "Some big ones",
			trunks:   []int{3, 35, 20, 4, 8, 6, 7},
			expected: 29,
		},
		{
			name:     "A lot of trunks",
			trunks:   []int{3, 35, 20, 4, 8, 6, 7, 3, 35, 20, 4, 40, 6, 7, 3, 2, 9, 12},
			expected: 68,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			income := calculateIncome(tt.trunks)
			if income != tt.expected {
				t.Errorf("wanted: %v, got: %v", tt.expected, income)
			}
		})
	}
}

func TestMain_calculateBestOrders(t *testing.T) {
	tests := []struct {
		name                  string
		test                  SawMillTest
		expectedIncome        int
		expectedOrderedTrunks [][]int
	}{
		{
			name:                  "Just a few small trunks",
			test:                  SawMillTest{TreeTrunks: []int{1, 2, 1}},
			expectedIncome:        3,
			expectedOrderedTrunks: [][]int{{1, 2, 1}},
		},
		{
			name:                  "Some big ones",
			test:                  SawMillTest{TreeTrunks: []int{3, 35, 20}},
			expectedIncome:        19,
			expectedOrderedTrunks: [][]int{{20, 35, 3}, {35, 20, 3}},
		},
		{
			name:                  "A lot of trunks",
			test:                  SawMillTest{TreeTrunks: []int{3, 35, 20, 4, 3, 6}},
			expectedIncome:        31,
			expectedOrderedTrunks: [][]int{{4, 3, 6, 3, 35, 20}, {4, 6, 3, 3, 35, 20}, {4, 3, 3, 6, 35, 20}, {4, 3, 3, 6, 20, 35}, {4, 3, 6, 3, 20, 35}, {4, 6, 3, 3, 20, 35}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			income, orderedTrunks := calculateBestOrders(tt.test)
			if income != tt.expectedIncome {
				t.Errorf("Wrong Income: wanted: %v, got: %v", tt.expectedIncome, income)
			}
			if !equalArrayofArrays(orderedTrunks, tt.expectedOrderedTrunks) {
				t.Errorf("Wrong ordered trunks: wanted: %v, got: %v", tt.expectedOrderedTrunks, orderedTrunks)
			}
		})
	}
}

func equalArrayofArrays(arrayOfArrays1 [][]int, arrayOfArrays2 [][]int) bool {
	if len(arrayOfArrays1) != len(arrayOfArrays2) {
		return false
	}
	for i, array1Elem := range arrayOfArrays1 {
		array2Elem := arrayOfArrays2[i]
		if len(array1Elem) != len(array2Elem) {
			return false
		}
		for j, elem := range array1Elem {
			if array2Elem[j] != elem {
				return false
			}
		}
	}
	return true
}
