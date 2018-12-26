package assertord

import "testing"

func TestIsOrdered(t *testing.T) {
	mockT := new(testing.T)

	arr := []int{1, 2, 3, 4}
	arr1 := []int{4, 3, 2, 1}

	if !IsOrdered(mockT, arr) {
		t.Error("IsOrdered should return true: arr is ordered")
	}
	if !IsOrdered(mockT, arr1) {
		t.Error("IsOrdered should return true: arr1 is ordered")
	}
}

func TestIsNotOrdered(t *testing.T) {
	mockT := new(testing.T)

	arr := []int{1, 2, 4, 3}
	if !IsNotOrdered(mockT, arr) {
		t.Error("IsNotOrdered should return true: array is not ordered")
	}
}

func TestIsOrderedAsc(t *testing.T) {
	mockT := new(testing.T)

	arr := []int{1, 2, 3, 4}
	arr1 := []int{4, 3, 2, 1}

	if !IsOrderedAsc(mockT, arr) {
		t.Error("IsOrderedAsc should return true: arr in asceding ordered")
	}
	if IsOrderedAsc(mockT, arr1) {
		t.Error("IsOrderedAsc should return false: arr1 not in asceding ordered")
	}
}

func TestIsNotOrderedAsc(t *testing.T) {
	mockT := new(testing.T)

	arr := []int{1, 2, 3, 4}
	arr1 := []int{4, 3, 2, 1}

	if IsNotOrderedAsc(mockT, arr) {
		t.Error("IsOrderedAsc should return false: arr in asceding ordered")
	}

	if !IsNotOrderedAsc(mockT, arr1) {
		t.Error("IsOrderedAsc should return true: arr1 is ordered")
	}
}

func TestIsOrderedDesc(t *testing.T) {
	mockT := new(testing.T)

	arr := []int{4, 3, 2, 1}
	arr1 := []int{1, 2, 3, 4}

	if !IsOrderedDesc(mockT, arr) {
		t.Error("IsOrderedDesc should return true: arr in asceding ordered")
	}
	if IsOrderedDesc(mockT, arr1) {
		t.Error("IsOrderedAsc should return false: arr1 not in asceding ordered")
	}
}

func TestIsNotOrderedDesc(t *testing.T) {
	mockT := new(testing.T)

	arr := []int{4, 3, 2, 1}
	arr1 := []int{1, 2, 3, 4}

	if IsNotOrderedDesc(mockT, arr) {
		t.Error("IsOrderedAsc should return false: arr in asceding ordered")
	}

	if !IsNotOrderedDesc(mockT, arr1) {
		t.Error("IsOrderedAsc should return true: arr1 is ordered")
	}
}
