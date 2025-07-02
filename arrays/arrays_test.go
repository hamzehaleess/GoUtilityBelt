package arrays

import (
        "reflect"
        "testing"
)

func TestContains(t *testing.T) {
        tests := []struct {
                slice    []int
                element  int
                expected bool
        }{
                {[]int{1, 2, 3, 4, 5}, 3, true},
                {[]int{1, 2, 3, 4, 5}, 6, false},
                {[]int{}, 1, false},
        }

        for _, test := range tests {
                result := Contains(test.slice, test.element)
                if result != test.expected {
                        t.Errorf("Contains(%v, %d) = %v; expected %v", test.slice, test.element, result, test.expected)
                }
        }
}

func TestRemove(t *testing.T) {
        tests := []struct {
                slice    []int
                element  int
                expected []int
        }{
                {[]int{1, 2, 3, 4, 5}, 3, []int{1, 2, 4, 5}},
                {[]int{1, 2, 3, 4, 5}, 6, []int{1, 2, 3, 4, 5}},
                {[]int{}, 1, []int{}},
        }

        for _, test := range tests {
                result := Remove(test.slice, test.element)
                if !reflect.DeepEqual(result, test.expected) {
                        t.Errorf("Remove(%v, %d) = %v; expected %v", test.slice, test.element, result, test.expected)
                }
        }
}

func TestUnique(t *testing.T) {
        tests := []struct {
                slice    []int
                expected []int
        }{
                {[]int{1, 2, 2, 3, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
                {[]int{1, 1, 1, 1}, []int{1}},
                {[]int{}, []int{}},
        }

        for _, test := range tests {
                result := Unique(test.slice)
                if !reflect.DeepEqual(result, test.expected) {
                        t.Errorf("Unique(%v) = %v; expected %v", test.slice, result, test.expected)
                }
        }
}

func TestSum(t *testing.T) {
        tests := []struct {
                slice    []int
                expected int
        }{
                {[]int{1, 2, 3, 4, 5}, 15},
                {[]int{-1, -2, -3}, -6},
                {[]int{}, 0},
        }

        for _, test := range tests {
                result := Sum(test.slice)
                if result != test.expected {
                        t.Errorf("Sum(%v) = %d; expected %d", test.slice, result, test.expected)
                }
        }
}

func TestMax(t *testing.T) {
        tests := []struct {
                slice    []int
                expected int
        }{
                {[]int{1, 2, 3, 4, 5}, 5},
                {[]int{-1, -2, -3}, -1},
                {[]int{42}, 42},
        }

        for _, test := range tests {
                result := Max(test.slice)
                if result != test.expected {
                        t.Errorf("Max(%v) = %d; expected %d", test.slice, result, test.expected)
                }
        }
}

func TestMin(t *testing.T) {
        tests := []struct {
                slice    []int
                expected int
        }{
                {[]int{1, 2, 3, 4, 5}, 1},
                {[]int{-1, -2, -3}, -3},
                {[]int{42}, 42},
        }

        for _, test := range tests {
                result := Min(test.slice)
                if result != test.expected {
                        t.Errorf("Min(%v) = %d; expected %d", test.slice, result, test.expected)
                }
        }
}

func TestReverse(t *testing.T) {
        tests := []struct {
                slice    []int
                expected []int
        }{
                {[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
                {[]int{1}, []int{1}},
                {[]int{}, []int{}},
        }

        for _, test := range tests {
                result := Reverse(test.slice)
                if !reflect.DeepEqual(result, test.expected) {
                        t.Errorf("Reverse(%v) = %v; expected %v", test.slice, result, test.expected)
                }
        }
}

func TestChunk(t *testing.T) {
        tests := []struct {
                slice    []int
                size     int
                expected [][]int
        }{
                {[]int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
                {[]int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}},
                {[]int{}, 2, nil},
        }

        for _, test := range tests {
                result := Chunk(test.slice, test.size)
                if !reflect.DeepEqual(result, test.expected) {
                        t.Errorf("Chunk(%v, %d) = %v; expected %v", test.slice, test.size, result, test.expected)
                }
        }
}

func TestFlatten(t *testing.T) {
        tests := []struct {
                slice    [][]int
                expected []int
        }{
                {[][]int{{1, 2}, {3, 4}, {5}}, []int{1, 2, 3, 4, 5}},
                {[][]int{{1}, {2}, {3}}, []int{1, 2, 3}},
                {[][]int{}, nil},
        }

        for _, test := range tests {
                result := Flatten(test.slice)
                if !reflect.DeepEqual(result, test.expected) {
                        t.Errorf("Flatten(%v) = %v; expected %v", test.slice, result, test.expected)
                }
        }
}

func TestShuffle(t *testing.T) {
        original := []int{1, 2, 3, 4, 5}
        result := Shuffle(original)
        
        // Test that original slice is not modified
        expected := []int{1, 2, 3, 4, 5}
        if !reflect.DeepEqual(original, expected) {
                t.Errorf("Shuffle modified original slice: got %v, expected %v", original, expected)
        }
        
        // Test that result has same length
        if len(result) != len(original) {
                t.Errorf("Shuffle result length = %d; expected %d", len(result), len(original))
        }
        
        // Test that result contains all original elements
        for _, elem := range original {
                if !Contains(result, elem) {
                        t.Errorf("Shuffle result missing element %d", elem)
                }
        }
        
        // Test empty slice
        empty := []int{}
        shuffledEmpty := Shuffle(empty)
        if len(shuffledEmpty) != 0 {
                t.Errorf("Shuffle of empty slice should return empty slice")
        }
}
