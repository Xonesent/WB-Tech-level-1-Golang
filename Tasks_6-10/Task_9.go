/*

№ 9
	Разработать конвейер чисел.
	Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	massive := [5]int{2, 4, 6, 8, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup

	for _, value := range massive {
		wg.Add(2)
		go write_num(ch1, &wg, value)
		go write_square(ch1, ch2, &wg)
	}

	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
	}()

	for value := range ch2 {
		fmt.Print(value, " ")
	}
}

func write_num(ch1 chan int, wg *sync.WaitGroup, value int) {
	defer wg.Done()
	ch1 <- value
}

func write_square(ch1, ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	value, ok := <-ch1
	if !ok {
		return
	}
	ch2 <- value * value
}
