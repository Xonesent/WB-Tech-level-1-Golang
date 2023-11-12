/*

№ 18 Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type Counter struct {
	num int
	mu  sync.Mutex
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("До какого значения будет счетчик?")
	var result int
	fmt.Fscan(reader, &result)

	counter := new(Counter)
	var wg sync.WaitGroup

	for i := 0; i < result; i++ {
		wg.Add(1)
		go Count(counter, &wg)
	}

	wg.Wait()
	fmt.Println(counter.num)
}

func Count(counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	counter.mu.Lock()
	counter.num++
	counter.mu.Unlock()
}
