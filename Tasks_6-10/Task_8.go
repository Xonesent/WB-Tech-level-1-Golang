/*

№ 8 Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func setBit(value int64, position uint, bitValue int) int64 {
	if bitValue == 1 {
		return value | (1 << position)
	} else if bitValue == 0 {
		return value &^ (1 << position)
	}
	return value
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var value int64
	var position uint
	var bitValue int

	fmt.Print("Введите число (int64): ")
	fmt.Fscan(reader, &value)

	fmt.Print("Введите номер бита (0-63): ")
	fmt.Fscan(reader, &position)

	fmt.Print("Установить бит в 1 или 0? (1/0): ")
	fmt.Fscan(reader, &bitValue)

	result := setBit(value, position, bitValue)
	fmt.Printf("Результат: %d\n", result)
}
