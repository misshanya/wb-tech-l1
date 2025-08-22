package main

import "fmt"

func main() {
	nums := []int{3, 1, 15, 17, 12, 20, 1, 4, 2}
	sorted := quickSort(nums)

	fmt.Println(sorted)
}

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	pivotIdx := len(nums) / 2
	pivot := nums[pivotIdx]

	left := []int{}
	middle := []int{}
	right := []int{}

	for _, num := range nums {
		if num < pivot {
			left = append(left, num)
		} else if num > pivot {
			right = append(right, num)
		} else {
			middle = append(middle, num)
		}
	}

	return append(append(quickSort(left), middle...), quickSort(right)...)
}
