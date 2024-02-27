package main

import "fmt"

func main() {
	fmt.Println("studing abt pointers")
	myNumb := 50
	var a = &myNumb
	fmt.Println("Address stored of a:", a)
	fmt.Println("Value stored in a:", *a)

}
