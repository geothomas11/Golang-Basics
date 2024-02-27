package main

// import (
// 	"fmt"
// 	"time"
// )

// func foo(s string) {
// 	for i := 1; i <= 3; i++ {
// 		fmt.Println(s, "", i)
// 		time.Sleep(100 * time.Millisecond)
// 	}

// }

// func main() {
// 	fmt.Println("Main started..")

// 	//Starting Two go routines
// 	go foo("1st goroutine :")
// 	go foo("2nd goroutine :")

// 	//Wait for the goroutine to finish before main goroutine ends
// 	time.Sleep(time.Second)
// 	fmt.Println("Main finished")
// }

// import "fmt"

// func main() {
// 	var c chan int
// 	if c == nil {
// 		fmt.Println("C is nil")
// 		fmt.Printf("type of c is %T\n", c)
// 		d := make(chan string)
// 		fmt.Printf("type of c is %T\n", d)
// 	}

// }

// func sendValues(myIntChannel chan int) {
// 	for i := 0; i < 5; i++ {
// 		myIntChannel <- i
// 	}

//}
