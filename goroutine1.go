package main

import (
	"fmt"
	"time"
	"math/rand"
)

const maxRoutine = 10
const maxParallelProcess = 1

func main() {
	ch := make(chan int)
	semaphore := make(chan int, maxParallelProcess)
	defer close(ch)
	defer close(semaphore)

	for i := 0; i < maxRoutine ; i++ {
		go func (s chan <- int, sem chan int, no int) {
			// 自分の番が来るのを待つ
			sem <- 0
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			s <- no
			<- sem
			fmt.Println("end")
		}(ch, semaphore, i)
	}

	fmt.Println("created!")

	// 受信用処理
	for i := 0; i < maxRoutine ; i++ {
		time.Sleep(1)
		value, _ := <-ch
		fmt.Println(value)
	}
	fmt.Println("all end!")
}
