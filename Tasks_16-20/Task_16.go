/*

№ 16 Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 3, 0, 90, 8, 5, 3, 3, 3, 1, 22}
	quick_sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quick_sort(arr []int, start, end int) {
	if start < end {
		left := start
		right := end
		mid := arr[(left+end)/2]

		for left < right {
			for arr[left] < mid {
				left++
			}
			for arr[right] > mid {
				right--
			}
			if left <= right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}

		quick_sort(arr, start, right)
		quick_sort(arr, left, end)
	}
}
