package main

import "fmt"

func main() {
	nums := []int{1, 2, 5, 7, 11, 12}
	fmt.Println(binSearch(nums, 7))
}

func binSearch(nums []int, x int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == x {
			return mid
		}
		if nums[mid] < x {
			l = mid + 1
		} else if nums[mid] > x {
			r = mid - 1
		}
	}

	return -1
}
