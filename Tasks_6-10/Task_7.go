/*

№ 7 Реализовать конкурентную запись данных в map.

*/

package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	map1 := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	var q_goroutines int
	fmt.Print("Выбери сколько горутин будет записывать значения в мапу\n")
	fmt.Scan(&q_goroutines)

	for i := 1; i <= q_goroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			map1[strconv.Itoa(i)] = i
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(map1)
}
