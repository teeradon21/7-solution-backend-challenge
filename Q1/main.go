package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	filePath := "../files/hard.json"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var array [][]int

	err = json.Unmarshal(bytes, &array)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	for line := len(array) - 2; line >= 0; line-- {
		for pos := 0; pos < len(array[line]); pos++ {
			array[line][pos] += max(array[line+1][pos], array[line+1][pos+1])
		}
	}

	fmt.Println(array[0][0])
}
