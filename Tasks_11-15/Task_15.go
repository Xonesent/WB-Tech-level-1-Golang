/*

№ 15
	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

*/

package main

import (
	"bytes"
	"fmt"
)

/*
	var justString string					<--- Глобальная переменная = не круто

	func someFunc() {
		v := createHugeString(1 << 10)		<--- Может выделить супер много лишней памяти, сборщик мусора не всегда мгновенно реагирует на такие ситуации
		justString = v[:100]				<--- Если строка v будет размером меньше 100 = рантайм error
	}

	func main() {
		someFunc()
	}
*/

func createHugeString(length int) string {
	var buf bytes.Buffer
	for i := 0; i < length; i++ {
		buf.WriteRune('A')
	}
	return buf.String()
}

func someFunc() string {
	v := createHugeString(1 << 10)
	var justString string
	if len(v) >= 100 {
		justString = v[:100]
	}
	return justString
}

func main() {
	fmt.Println(someFunc())
}
