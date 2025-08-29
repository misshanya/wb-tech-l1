package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	idxToDelete := 2

	copy(nums[idxToDelete:], nums[idxToDelete+1:])
	nums = nums[:len(nums)-1]
	fmt.Println(nums)
}
