package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func () {
		for i := 0; i < 10; i++ {
			select {
			case ch1 <- 0:
				fmt.Println("Send to ch1!!")
			case ch2 <- "test":
				fmt.Println("Send to ch2!!")
			}
		}
		close(ch1)
		close(ch2)
	} ()

	for {
		var ok1, ok2 bool
		var val1 int
		var val2 string

		select {
		case val1, ok1 = <- ch1:
			fmt.Println(val1)
		case val2, ok2 = <- ch2:
			fmt.Println(val2)
		}
		if !ok1 && !ok2 {
			break
		}
	}
	fmt.Println("all end!")
}
