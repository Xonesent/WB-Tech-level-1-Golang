/*

№ 11 Реализовать пересечение двух неупорядоченных множеств.

*/

package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 0}

	fmt.Println(intersaction(slice1, slice2))
}

func intersaction(slice1, slice2 []int) []int {
	anslist := make([]int, 0)
	ansmap := make(map[int]struct{})

	for i := 0; i < len(slice1); i++ {
		ansmap[slice1[i]] = struct{}{}
	}

	for i := 0; i < len(slice2); i++ {
		if _, ok := ansmap[slice2[i]]; ok {
			anslist = append(anslist, slice2[i])
		}
	}

	return anslist
}
