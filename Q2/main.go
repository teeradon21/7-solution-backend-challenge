package main

import (
	"fmt"
	"strconv"
)

func findDecrease(encode string, index int) int {
	decrease := 0
	for i := index; i < len(encode); i++ {
		if encode[i] == 'L' {
			decrease++
		} else if encode[i] == 'R' {
			return decrease
		}
	}
	return decrease
}

func findEqualBackward(encode string, index int) int {
	equal := 0
	for i := index - 1; i >= 0; i-- {
		if encode[i] == '=' {
			equal++
		} else {
			return equal
		}
	}
	return equal
}

func main() {

	var encoded string
	fmt.Scan(&encoded)
	numbers := make([]int, len(encoded)+1)
	numbers[0] = 0
	decreasing := false
	for i := 0; i < len(encoded); i++ {
		if encoded[i] == 'L' {
			if !decreasing {
				decrease := findDecrease(encoded, i)
				if decrease > numbers[i] {
					equal := findEqualBackward(encoded, i)
					for j := i; j >= i-equal; j-- {
						numbers[j] = decrease
					}
				}
				numbers[i+1] = decrease - 1
				decreasing = true
			} else {
				numbers[i+1] = numbers[i] - 1
			}
		} else if encoded[i] == 'R' {
			numbers[i+1] = numbers[i] + 1
			decreasing = false
		} else if encoded[i] == '=' {
			numbers[i+1] = numbers[i]
		}
	}
	sol := ""
	for i := 0; i < len(numbers); i++ {
		sol += strconv.Itoa(numbers[i])
	}
	fmt.Println(sol)
}
