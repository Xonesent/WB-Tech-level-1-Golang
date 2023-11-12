/*

№ 26
	Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
	Функция проверки должна быть регистронезависимой.

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

	fmt.Println("Введите слово для проверки")
	var word string
	fmt.Fscan(reader, &word)

	fmt.Println(checker(word))
}

func checker(str string) bool {
	str = strings.ToLower(str)
	map1 := make(map[rune]struct{})

	for _, value := range []rune(str) {
		_, ok := map1[value]
		if ok {
			return false
		} else {
			map1[value] = struct{}{}
		}
	}

	return true
}
