/*

№ 14 Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

*/

package main

import (
	"fmt"
)

func main() {
	var val interface{}
	val = make(chan interface{})

	switch v := val.(type) {
	case int:
		fmt.Printf("Это целое число: %d\n", v)
	case string:
		fmt.Printf("Это строка: %s\n", v)
	case bool:
		fmt.Printf("Это булево значение: %t\n", v)
	case chan interface{}:
		fmt.Println("Это канал")
	default:
		fmt.Println("Не удалось определить тип")
	}
}
