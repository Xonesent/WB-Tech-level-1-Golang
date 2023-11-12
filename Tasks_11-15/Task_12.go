/*

№ 12 Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

*/

package main

import (
	"fmt"
)

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}

	map1 := make(map[string]struct{})
	for _, value := range slice {
		map1[value] = struct{}{}
	}

	anslist := make([]string, 0)
	for key := range map1 {
		anslist = append(anslist, key)
	}

	fmt.Println(anslist)
}
