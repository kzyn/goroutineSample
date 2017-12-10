package main

import (
	"fmt"
	"os"
	"runtime"
)

const maxRoutine = 10

func main () {
	ch := make(chan int)
	defer close(ch)

	fmt.Printf("cpu:%v\n", runtime.NumCPU())
	fmt.Printf("maxgoroutine:%v\n", runtime.GOMAXPROCS(0))

	for i := 0; i < maxRoutine; i++ {
		go func () {
			val, _ := <-ch
			val++
			fmt.Println(val)
			if val == maxRoutine {
				os.Exit(0)
			}
			ch <- val
		} ()
	}
	// 初期値を投入
	ch <- 0
	for {
	}
}
