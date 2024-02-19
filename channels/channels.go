package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels")

	mycha := make(chan int, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)

	//Recive Only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-mycha
		fmt.Println(isChannelOpen)
		fmt.Println(val)

		wg.Done()

	}(mycha, &wg)

	go func(ch chan<- int, wq *sync.WaitGroup) {
		mycha <- 5
		close(mycha)

		wg.Done()
	}(mycha, &wg)

	wg.Wait()

}
