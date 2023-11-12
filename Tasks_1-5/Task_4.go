/*

№ 4
	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.
	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

*/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func channel_write(ch chan string) {
	for {
		fmt.Println("Что записать в канал?")
		var value string
		fmt.Scan(&value)
		ch <- value
	}
}

func channel_read(ch chan string, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for {
		value, ok := <-ch
		if !ok {
			fmt.Printf("Воркер %d завершил работу.\n", i)
			return
		}
		fmt.Printf("Рабочий %d прочитал %s\n", i, value)
	}
}

func programm_end(sig chan os.Signal, ch chan string, wg *sync.WaitGroup) {
	<-sig
	close(ch)
	wg.Wait()

	fmt.Println("Программа завершила работу.")
	os.Exit(0)
}

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	var q_workers int
	fmt.Println("Введи количество рабочих")
	fmt.Scan(&q_workers)

	go channel_write(ch)
	for i := 1; i <= q_workers; i++ {
		wg.Add(1)
		go channel_read(ch, &wg, i)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	go programm_end(sig, ch, &wg)
	select {}
}
