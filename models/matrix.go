package models

import (
	"strconv"
	"strings"
)

// FormatRecords formats the records into a string in matrix format
func FormatRecords(records [][]string) string {
	var response strings.Builder
	for _, row := range records {
		response.WriteString(strings.Join(row, ","))
		response.WriteString("\n")
	}
	return response.String()
}

// InvertMatrix inverts the matrix
func InvertMatrix(records [][]string) [][]string {
	if len(records) == 0 {
		return records
	}

	inverted := make([][]string, len(records[0]))
	for i := range inverted {
		inverted[i] = make([]string, len(records))
	}

	for i, row := range records {
		for j, value := range row {
			inverted[j][i] = value
		}
	}
	return inverted
}

// FlattenMatrix flattens the matrix into a single line
func FlattenMatrix(records [][]string) string {
	flat := make([]string, 0, len(records)*len(records[0]))
	for _, row := range records {
		flat = append(flat, row...)
	}
	return strings.Join(flat, ",")
}

// SumMatrix calculates the sum of the integers in the matrix
func SumMatrix(records [][]string) int {
	sum := 0
	for _, row := range records {
		for _, value := range row {
			if num, err := strconv.Atoi(value); err == nil {
				sum += num
			}
		}
	}
	return sum
}

// MultiplyMatrix calculates the product of the integers in the matrix
func MultiplyMatrix(records [][]string) int {
	product := 1
	hasNonZero := false
	for _, row := range records {
		for _, value := range row {
			if num, err := strconv.Atoi(value); err == nil {
				product *= num
				hasNonZero = true
			}
		}
	}
	if !hasNonZero {
		return 0 // Return 0 if no valid numbers were found
	}
	return product
}
