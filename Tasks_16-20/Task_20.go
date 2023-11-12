/*

№ 20 Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Напиши предложение, которую мне перевернуть")
	var strs []string

	line, _, _ := reader.ReadLine()
	strs = strings.Fields(string(line))

	fmt.Println(reverse(strs))
}

func reverse(strs []string) []string {
	for i := 0; i < len(strs)/2; i++ {
		strs[i], strs[len(strs)-1-i] = strs[len(strs)-1-i], strs[i]
	}
	return strs
}
