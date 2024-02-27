package main

import "fmt"

const (
	car = iota //"IOTA"it means that it will autoincrement
	bus
	train
)

func main() {
	fmt.Println(car)
	fmt.Println(bus)
	fmt.Println(train)

}
