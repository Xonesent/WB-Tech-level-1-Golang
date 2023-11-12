/*

№ 19 Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Выбери строку, которую мне перевернуть")
	var str string
	fmt.Fscan(reader, &str)

	fmt.Println(reverse(str))
}

func reverse(str string) string {
	tmp := []byte(str)
	for i := 0; i < len(str)/2; i++ {
		tmp[i], tmp[len(str)-1-i] = tmp[len(str)-1-i], tmp[i]
	}
	return string(tmp)
}
