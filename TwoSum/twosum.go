package main

import "fmt"

func twoSum(nums []int, target int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}

		}
	}
	return nil
}

func main() {
	target := 9
	numbers := []int{1, 2, 4, 5, 5}
	fmt.Println(twoSum(numbers, target))
}
