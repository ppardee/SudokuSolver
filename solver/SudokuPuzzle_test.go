package solver

import "testing"

func TestContains(t *testing.T) {
	testData := []int{1, 2, 3, 4, 5}
	if contains(testData, 0) {
		t.Errorf("Array does not contain 0")
	}

	if !contains(testData, 1) {
		t.Errorf("Array does contain 1")
	}
}

func TestIntToBits(t *testing.T) {
	testData := map[int]int{
		0: 0,
		1: 0b000000001,
		2: 0b000000010,
		3: 0b000000100,
		4: 0b000001000,
		5: 0b000010000,
		6: 0b000100000,
		7: 0b001000000,
		8: 0b010000000,
		9: 0b100000000,
	}

	for key, value := range testData {
		if res := intToBits(key); res != value {
			t.Errorf("Expected %09b but got %09b", value, res)
		}
	}

}

func TestBitsToInts(t *testing.T) {
	testData := map[int][]int{
		0b100000001: []int{1, 9},
		0b111111111: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		0b010000011: []int{1, 2, 8},
		0b000111000: []int{4, 5, 6},
	}

	for key, value := range testData {
		if res := bitsToInts(key); !slicesAreEquivalent(res, value) {
			t.Errorf("Expected %v but got %v", value, res)
		}
	}

}

func slicesAreEquivalent(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for _, v := range a {
		if !contains(a, v) {
			return false
		}
	}
	return true

}

func TestGetNonantStart(t *testing.T) {
	testData := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 3,
		5: 3,
		6: 3,
		7: 6,
		8: 6,
		9: 6,
	}
	for key, val := range testData {
		if res := getNonantStart(key); res != val {
			t.Errorf("Expected %v but got %v", val, res)
		}
	}

}
