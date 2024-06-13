package main

import (
	"bufio"
	"os"
	"strconv"
)

func ReadDataFromFile(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		data = append(data, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
