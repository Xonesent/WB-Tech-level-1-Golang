/*

№ 10
	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
	Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	map1 := make(map[int][]float64)

	for {
		fmt.Print("Введите температуру\n")
		var value float64
		_, err := fmt.Fscan(reader, &value)
		if err != nil {
			break
		}
		map1[int(value/10)*10] = append(map1[int(value/10)*10], value)
	}

	var count int = 0
	for key, value := range map1 {
		fmt.Printf("%d:{", key)
		for i := 0; i < len(value); i++ {
			if i != len(value)-1 {
				fmt.Printf("%.1f, ", value[i])
			} else {
				fmt.Printf("%.1f}", value[i])
			}
		}
		if count != len(map1)-1 {
			fmt.Print(", ")
		}
		count++
	}
}
