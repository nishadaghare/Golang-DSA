package main

import "fmt"

func distinct(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, val := range arr {
		if !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 4, 5}
	fmt.Println("Distinct elements:", distinct(arr))
}
