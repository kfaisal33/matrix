package models

import (
	"testing"
)

func TestFormatRecords(t *testing.T) {
	records := [][]string{{"1", "2", "3"}, {"4", "5", "6"}}
	expected := "1,2,3\n4,5,6\n"
	if got := FormatRecords(records); got != expected {
		t.Errorf("FormatRecords() = %q; want %q", got, expected)
	}
}

func TestInvertMatrix(t *testing.T) {
	tests := []struct {
		records  [][]string
		expected [][]string
	}{
		{
			records:  [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			expected: [][]string{{"1", "4"}, {"2", "5"}, {"3", "6"}},
		},
	}

	for _, test := range tests {
		if got := InvertMatrix(test.records); !equal(got, test.expected) {
			t.Errorf("InvertMatrix(%v) = %v; want %v", test.records, got, test.expected)
		}
	}
}

func TestFlattenMatrix(t *testing.T) {
	records := [][]string{{"1", "2"}, {"3", "4"}}
	expected := "1,2,3,4"
	if got := FlattenMatrix(records); got != expected {
		t.Errorf("FlattenMatrix() = %q; want %q", got, expected)
	}
}

func TestSumMatrix(t *testing.T) {
	records := [][]string{{"1", "2"}, {"3", "4"}}
	expected := 10
	if got := SumMatrix(records); got != expected {
		t.Errorf("SumMatrix() = %d; want %d", got, expected)
	}
}

func TestMultiplyMatrix(t *testing.T) {
	records := [][]string{{"1", "2"}, {"3", "4"}}
	expected := 24 // 1*2*3*4
	if got := MultiplyMatrix(records); got != expected {
		t.Errorf("MultiplyMatrix() = %d; want %d", got, expected)
	}
}

// Helper function to compare two 2D slices
func equal(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
