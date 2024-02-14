package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Open the file for reading
	os := os.Args[1:]
	if len(os) == 0 {
		fmt.Println("No file name!!")
		return
	}

	filePath := os[0]
	content, err := readFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	if len(content) == 0 {
		fmt.Println("Empty")
		return
	}
	fmt.Printf("Average: %v\n", int(math.Round(avg(content))))
	fmt.Printf("Median: %v\n", int(math.Round(med(content))))
	fmt.Printf("Variance: %v\n", int(math.Round(variance(content))))
	fmt.Printf("Standard Deviation: %v\n", int(math.Round(sd(content))))
}

func avg(arr []float64) float64 {
	var cnt float64
	for i := 0; i < len(arr); i++ {
		cnt = cnt + arr[i]
	}
	var v float64
	v = cnt / float64(len(arr))
	return v
}

func med(arr []float64) float64 {
	sort.Float64s(arr)
	mid := len(arr) / 2
	if len(arr)%2 == 0 {
		return (arr[mid] + arr[mid-1]) / 2
	} else {
		return arr[mid]
	}
}

func variance(arr []float64) float64 {
	var res float64
	var cnt float64

	for i := 0; i < len(arr); i++ {
		cnt += (arr[i] - avg(arr)) * (arr[i] - avg(arr))
	}
	res = cnt / float64(len(arr))
	return res
}

func sd(arr []float64) float64 {
	return math.Sqrt(variance(arr))
}

func readFile(filePath string) ([]float64, error) {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Close the file when the function exits

	// Read the entire content of the file
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []float64

	for fileScanner.Scan() {
		number, err := strconv.ParseFloat(fileScanner.Text(), 64)
		if err != nil {
			fmt.Println("Invalid data")
			return nil, err
		}
		fileLines = append(fileLines, number)

	}

	return fileLines, nil
}
