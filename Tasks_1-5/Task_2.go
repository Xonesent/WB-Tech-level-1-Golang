/*

№ 2 Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	massive := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, value := range massive {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Print(num*num, " ")
		}(value)
	}

	wg.Wait()
}
