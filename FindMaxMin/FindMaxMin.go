package main

import "fmt"

func main() {
	numbers := []int{22, 44, 23, 65, 18, 74, 28}
	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	fmt.Println(numbers)
	fmt.Println(min)
	fmt.Println(max)

}
