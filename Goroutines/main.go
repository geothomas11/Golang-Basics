// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 3; i++ {
// 		wg.Add(1)
// 		go func(num int) {
// 			defer wg.Done()
// 			fmt.Printf("Goroutine %d: Hello World\n", num)
// 		}(i)
// 	}

// 	wg.Wait()

// }

package main

import "fmt"

func sendValues(myChannel chan int) {
	for i := 0; i <= 5; i++ {
		myChannel <- i
	}

}

func main() {
	myChannel := make(chan int)
	go sendValues(myChannel)

	for i := 0; i <= 5; i++ {
		fmt.Println(<-myChannel)
	}
}
