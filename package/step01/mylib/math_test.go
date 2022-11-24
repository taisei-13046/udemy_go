package mylib

import "testing"

// TextXxxxという命名
// fileName_test.go
func TestAverage(t *testing.T) {
	v := Average([]int{1, 2, 3, 4, 5})

	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}
