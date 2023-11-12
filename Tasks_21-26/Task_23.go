/*

№ 23 Удалить i-ый элемент из слайса.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Выберит индекс элемента для его удаления из слайса")
	var index int
	fmt.Fscan(reader, &index)

	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(remove(slice, index))
}

func remove(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		fmt.Println(":(")
		return slice
	}
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}
