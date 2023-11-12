/*

№ 5
	Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
	По истечению N секунд программа должна завершаться.

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	fmt.Print("Выберите время секунд действия программы\n")
	var q_seconds int
	fmt.Scan(&q_seconds)

	timer := time.NewTimer(time.Duration(q_seconds) * time.Second)

	go func() {
		for {
			fmt.Print("Выберите что записать в канал\n")
			var data string
			fmt.Scan(&data)

			ch <- data
		}
	}()

	for {
		select {
		case value := <-ch:
			fmt.Print(value, " ")
		case <-timer.C:
			fmt.Print("Время вышло\n")
			return
		}
	}
}
