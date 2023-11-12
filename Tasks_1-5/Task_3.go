/*

№ 3 Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var sum int
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sum += n * n
		}(num)
	}
	wg.Wait()

	fmt.Print(sum)
}
