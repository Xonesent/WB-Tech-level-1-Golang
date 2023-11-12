/*

№ 6.3 Реализовать все возможные способы остановки выполнения горутины.

*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(go1 context.Context, wg *sync.WaitGroup, i int) {
	wg.Add(1)
	for {
		select {
		case <-go1.Done():
			wg.Done()
			fmt.Printf("Горутина %d завершена.\n", i)
			return
		default:
			fmt.Printf("Горутина %d выполняет работу...\n", i)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	go1, stop1 := context.WithCancel(context.Background())
	go2, stop2 := context.WithCancel(context.Background())

	go worker(go1, &wg, 1)
	go worker(go2, &wg, 2)

	time.Sleep(5 * time.Second)
	stop1()
	time.Sleep(2 * time.Second)
	stop2()
	time.Sleep(2 * time.Second)

	wg.Wait()
}
