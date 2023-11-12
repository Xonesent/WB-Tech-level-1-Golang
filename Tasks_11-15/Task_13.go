/*

№ 13 Поменять местами два числа без создания временной переменной.

*/

package main

import (
	"fmt"
)

func main() {
	var a, b int = 1, 2

	a, b = b, a
	fmt.Println(a, b)
}
