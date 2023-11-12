/*

№ 25 Реализовать собственную функцию sleep.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите время ожидания в секундах")
	var wait_time float64
	fmt.Fscan(reader, &wait_time)

	sleep(wait_time)
}

func sleep(wait_time float64) {
	t := time.Now()
	for time.Since(t).Seconds() < wait_time {
	}
}
