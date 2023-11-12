/*

№ 6.2 Реализовать все возможные способы остановки выполнения горутины.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func end_goroutine(stop chan string, wg *sync.WaitGroup) {
	for {
		select {
		case <-stop:
			fmt.Println("Горутина завершена.")
			wg.Done()
			return
		default:
			fmt.Println("Выполняю работу...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	stop := make(chan string)

	wg.Add(1)
	go end_goroutine(stop, &wg)

	time.Sleep(5 * time.Second)
	stop <- "1"

	wg.Wait()
}
