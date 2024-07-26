package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
)

type Range struct {
	Min    string `json:"min"`
	Max    string `json:"max"`
	Status string `json:"status"`
}

func main() {
	// Input range
	inputRange := Range{
		Min:    "0x80000",
		Max:    "0x7fffff",
		Status: "1HsMJxNiV7TLxmoF6uJNkydxPFDog4NQum, 14oFNXucftsHiUMY8uctg6N487riuyXs4h, 1CfZWK1QTQE3eS9qn61dQjV89KDjZzfNcv, 1L2GM8eE7mJWLdo3HZS6su1832NX2txaac",
	}

	// Number of intervals
	numIntervals := 50

	// Convert hex values to big integers
	min, _ := new(big.Int).SetString(inputRange.Min[2:], 16) // Remove "0x"
	max, _ := new(big.Int).SetString(inputRange.Max[2:], 16)

	// Calculate interval size
	intervalSize := new(big.Int).Sub(max, min)
	intervalSize.Div(intervalSize, big.NewInt(int64(numIntervals)))

	// Generate ranges
	ranges := make([]Range, numIntervals)
	for i := 0; i < numIntervals; i++ {
		rangeMin := new(big.Int).Set(min)
		rangeMax := new(big.Int).Add(min, intervalSize)

		// Adjust last range to include the maximum value
		if i == numIntervals-1 {
			rangeMax = new(big.Int).Set(max)
		}

		ranges[i] = Range{
			Min:    "0x" + hex.EncodeToString(rangeMin.Bytes()),
			Max:    "0x" + hex.EncodeToString(rangeMax.Bytes()),
			Status: inputRange.Status,
		}

		min.Add(min, intervalSize) // Move to the next interval
	}

	// Create output data structure
	output := struct {
		Ranges []Range `json:"ranges"`
	}{
		Ranges: ranges,
	}

	// Create JSON file
	file, err := os.Create("ranges.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write JSON data to file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print
	if err := encoder.Encode(output); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}