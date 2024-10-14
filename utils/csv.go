// package utils

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"net/http"
// )

// // ParseCSVFile parses the uploaded CSV file and returns the records
// func ParseCSVFile(r *http.Request) ([][]string, error) {
// 	file, _, err := r.FormFile("file")
// 	if err != nil {
// 		return nil, fmt.Errorf("error reading file: %w", err)
// 	}
// 	defer file.Close()

//		reader := csv.NewReader(file)
//		records, err := reader.ReadAll()
//		if err != nil {
//			return nil, fmt.Errorf("error parsing CSV: %w", err)
//		}
//		return records, nil
//	}
package utils

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

// ParseCSVFile parses the uploaded CSV file and returns the records
func ParseCSVFile(r *http.Request) ([][]string, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("Error closing file: %v\n", cerr)
		}
	}()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error parsing CSV: %w", err)
	}

	return records, nil
}
