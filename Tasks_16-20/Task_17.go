/*

№ 17 Реализовать бинарный поиск встроенными методами языка.

*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 3, 3, 5, 7, 10}
	fmt.Println(lbinsearch(arr, 0), rbinsearch(arr, 0))
	fmt.Println(lbinsearch(arr, 2), rbinsearch(arr, 2))
	fmt.Println(lbinsearch(arr, 4), rbinsearch(arr, 4))
	fmt.Println(lbinsearch(arr, 11), rbinsearch(arr, 11))
}

func lbinsearch(arr []int, num int) int {
	left := 0
	right := len(arr) - 1
	for left < right {
		mid_ind := (left + right) / 2
		if num > arr[mid_ind] {
			left = mid_ind + 1
		} else {
			right = mid_ind
		}
	}
	return left
}

func rbinsearch(arr []int, num int) int {
	left := 0
	right := len(arr) - 1
	for left < right {
		mid_ind := (left + right + 1) / 2
		if num < arr[mid_ind] {
			right = mid_ind - 1
		} else {
			left = mid_ind
		}
	}
	return left
}
